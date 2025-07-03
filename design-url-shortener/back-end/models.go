package main

import (
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type HandlerAPI struct {
	DBClient    *gorm.DB
	RedisClient *redis.Client
}

type RequestData struct {
	URL string `json:"url"`
}

type URL struct {
	ExpiredTime time.Time
	LongURL     string
	ShortURL    string
	ID          int
}
