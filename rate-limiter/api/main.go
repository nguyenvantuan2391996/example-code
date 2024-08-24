package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

const MaxRequestOneMinute = 60

type HandlerAPI struct {
	RedisClient *redis.Client
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

func getIPFromRequest(r *http.Request) string {
	ips := r.Header.Get("X-Forwarded-For")
	ipList := strings.Split(ips, ",")
	for _, ip := range ipList {
		if ip = strings.TrimSpace(ip); ip != "" && ip != "::1" && ip != "127.0.0.1" {
			return ip
		}
	}

	ip := r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func (h *HandlerAPI) RateLimiter(next http.Handler) http.Handler {
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

		// rate limit
		ip := getIPFromRequest(r)
		script := `
		local currentCount = tonumber(redis.call('GET', KEYS[1]) or '0')
		if currentCount == 0 then
			redis.call('SET', KEYS[1], 0, 'EX', ARGV[1])
		end

		redis.call('SET', KEYS[1], currentCount + 1, 'KEEPTTL')

		if currentCount > tonumber(ARGV[2]) then
			return "not pass"
		else
			return "pass"
		end`

		// Running Lua Script
		resultStr, err := h.RedisClient.Eval(context.Background(), script, []string{ip}, 60, MaxRequestOneMinute).Result()
		if err != nil {
			logrus.Warnf("Running Lua Script is failed with err: %v", err)
			return
		}

		if resultStr == "not pass" {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *HandlerAPI) testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Hello Viet Nam")
	if err != nil {
		return
	}
}

func main() {
	redisClient, err := initRedis("redis://default:@localhost:6379")
	if err != nil {
		panic("failed to init redis")
	}

	handler := HandlerAPI{
		RedisClient: redisClient,
	}

	mux := http.NewServeMux()
	mux.Handle("/test", handler.RateLimiter(http.HandlerFunc(handler.testHandler)))

	// Start the server
	log.Fatal(http.ListenAndServe(":3000", mux))
}
