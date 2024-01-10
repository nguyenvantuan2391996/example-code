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

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (h *HandlerAPI) generateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
		tmp, errGenerate := GenerateURL()
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
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *HandlerAPI) Create(ctx context.Context, record *URL) error {
	err := h.DBClient.Create(record).Error
	if err != nil {
		return err
	}

	data, errMarshal := json.Marshal(record)
	if errMarshal != nil {
		return nil
	}
	_, err = h.RedisClient.Set(ctx, record.ShortURL, data, ExpiredHours*time.Hour).Result()
	if err != nil {
		return err
	}

	return nil
}

func (h *HandlerAPI) GetURLByQueries(ctx context.Context, queries map[string]interface{}) (*URL, error) {
	var record *URL

	// get from redis
	result, errRedis := h.RedisClient.Get(ctx, fmt.Sprintf("%s", queries["short_url"])).Result()
	if errRedis == nil {
		errUnmarshal := json.Unmarshal([]byte(result), &record)
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

	mux := http.NewServeMux()
	mux.Handle("/generate-short-url", Middleware(http.HandlerFunc(handler.generateShortUrlHandler)))
	mux.Handle("/", Middleware(http.HandlerFunc(handler.redirectHandler)))

	// Start the server
	log.Fatal(http.ListenAndServe(":3000", mux))
}
