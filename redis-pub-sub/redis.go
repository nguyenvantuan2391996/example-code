package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Redis *redis.Client
}

func (r *RedisClient) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return r.Redis.Subscribe(ctx, channel)
}

func (r *RedisClient) Publish(ctx context.Context, channel, message string) (int64, error) {
	return r.Redis.Publish(ctx, channel, message).Result()
}
