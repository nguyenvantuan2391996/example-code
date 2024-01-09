package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aidarkhanov/nanoid/v2"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Domain             = "http://localhost:3000/"
	NumberOfCharacters = 7
	MaxRetryGenerate   = 10
	ExpiredHours       = 24
	DefaultExpiredTime = 90 * 24 * time.Hour
	PageNotFound       = "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><title>404 Not Found</title><link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/css/bootstrap.min.css\"><link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css?family=Arvo\"><style>.page_404{padding:40px 0;background:#fff;font-family:Arvo,serif}.page_404 img{width:100%}.four_zero_four_bg{background-image:url(https://cdn.dribbble.com/users/285475/screenshots/2083086/dribbble_1.gif);height:400px;background-position:center}.four_zero_four_bg h1{font-size:80px}.four_zero_four_bg h3{font-size:80px}.link_404{color:#fff!important;padding:10px 20px;background:#39ac31;margin:20px 0;display:inline-block}.contant_box_404{margin-top:-50px}</style></head><body><section class=\"page_404\"><div class=\"container\"><div class=\"row\"><div class=\"col-sm-12\"><div class=\"col-sm-10 col-sm-offset-1 text-center\"><div class=\"four_zero_four_bg\"><h1 class=\"text-center\">404</h1></div><div class=\"contant_box_404\"><h3 class=\"h2\">Look like you're lost</h3><p>the page you are looking for not avaible!</p><a href=\"\" class=\"link_404\">Go to Home</a></div></div></div></div></div></section></body></html>"
)

type HandlerAPI struct {
	DBClient    *gorm.DB
	RedisClient *redis.Client
}

type RequestData struct {
	URL string `json:"url"`
}

type URL struct {
	ID          int
	LongURL     string
	ShortURL    string
	ExpiredTime time.Time
}

func (h *HandlerAPI) generateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Read the body request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var request RequestData
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Error unmarshal request body", http.StatusBadRequest)
		return
	}

	// Generate short link by using nanoid
	shortLink := ""
	for i := 0; i < MaxRetryGenerate; i++ {
		tmp, errGenerate := generateURL()
		if errGenerate != nil {
			continue
		}

		_, errGet := h.GetURLByQueries(r.Context(), map[string]interface{}{
			"short_url": Domain + tmp,
		})
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			shortLink = Domain + tmp
			break
		}
	}

	if len(shortLink) == 0 {
		http.Error(w, "Error generate short url", http.StatusBadRequest)
		return
	}

	// create record
	err = h.Create(r.Context(), &URL{
		LongURL:     request.URL,
		ShortURL:    shortLink,
		ExpiredTime: time.Now().Add(DefaultExpiredTime),
	})
	if err != nil {
		http.Error(w, "Error generate short url", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, shortLink)
	if err != nil {
		return
	}
}

func (h *HandlerAPI) redirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	longURL, err := h.GetURLByQueries(r.Context(), map[string]interface{}{
		"short_url": Domain + strings.TrimLeft(r.RequestURI, "/"),
	})
	if err != nil || (longURL != nil && time.Now().After(longURL.ExpiredTime)) {
		w.Header().Set("Content-Type", "text/html")
		_, err := fmt.Fprint(w, PageNotFound)
		if err != nil {
			return
		}
		return
	}

	http.Redirect(w, r, longURL.LongURL, http.StatusMovedPermanently)
}

func (h *HandlerAPI) Create(ctx context.Context, record *URL) error {
	err := h.DBClient.Create(record).Error
	if err != nil {
		return err
	}

	go func() {
		h.RedisClient.Set(ctx, record.ShortURL, record.LongURL, ExpiredHours*time.Hour)
	}()

	return nil
}

func (h *HandlerAPI) GetURLByQueries(ctx context.Context, queries map[string]interface{}) (*URL, error) {
	var record *URL

	// get from redis
	list, errRedis := h.RedisClient.Get(ctx, fmt.Sprintf("%s", queries["short_url"])).Result()
	if errRedis == nil {
		errUnmarshal := json.Unmarshal([]byte(list), &record)
		if errUnmarshal != nil {
			return nil, errUnmarshal
		}
		return record, nil
	}

	err := h.DBClient.WithContext(ctx).Where(queries).First(&record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}

func generateURL() (string, error) {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	size := NumberOfCharacters

	key, err := nanoid.GenerateString(alphabet, size)
	if err != nil {
		return "", err
	}

	return key, nil
}

func initDatabase() (*gorm.DB, error) {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/shortner_url?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open connect to database:", err)
		return nil, nil
	}

	return db, nil
}

func initRedis(redisUrl string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatal("failed to connect redis:", err)
		return nil, nil
	}

	opts.PoolSize = 30
	opts.ReadTimeout = 5 * time.Second
	opts.WriteTimeout = 5 * time.Second
	opts.Username = ""

	redisClient := redis.NewClient(opts)

	cmd := redisClient.Ping(context.Background())
	if cmd.Err() != nil {
		log.Fatal("failed to ping redis: ", cmd.Err())
		return nil, nil
	}

	return redisClient, nil
}

func main() {
	dbClient, err := initDatabase()
	if err != nil {
		panic("failed to init database")
	}

	redisClient, err := initRedis("redis://default:admin123@localhost:6379")
	if err != nil {
		panic("failed to init redis")
	}

	handler := HandlerAPI{
		DBClient:    dbClient,
		RedisClient: redisClient,
	}

	http.HandleFunc("/generate-short-url", handler.generateShortUrlHandler)
	http.HandleFunc("/", handler.redirectHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":3000", nil))
}
