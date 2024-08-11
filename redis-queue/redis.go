package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Redis *redis.Client
}

func (r *RedisClient) LPush(ctx context.Context, queue string, value interface{}) (int64, error) {
	return r.Redis.LPush(ctx, queue, value).Result()
}

func (r *RedisClient) BRPop(ctx context.Context, queue string, timeout time.Duration) ([]string, error) {
	return r.Redis.BRPop(ctx, timeout, queue).Result()
}
