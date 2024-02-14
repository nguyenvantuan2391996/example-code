package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Redis *redis.Client
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expire time.Duration) (string, error) {
	return r.Redis.Set(ctx, key, value, expire).Result()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Redis.Get(ctx, key).Result()
}

func (r *RedisClient) Del(ctx context.Context, keys ...string) (int64, error) {
	return r.Redis.Del(ctx, keys...).Result()
}

func (r *RedisClient) Incr(ctx context.Context, key string) (int64, error) {
	return r.Redis.Incr(ctx, key).Result()
}

func (r *RedisClient) Decr(ctx context.Context, key string) (int64, error) {
	return r.Redis.Decr(ctx, key).Result()
}

func (r *RedisClient) Exists(ctx context.Context, keys ...string) (int64, error) {
	return r.Redis.Exists(ctx, keys...).Result()
}

// ZIncrBy Zset
func (r *RedisClient) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	return r.Redis.ZIncrBy(ctx, key, increment, member).Result()
}

func (r *RedisClient) ZScore(ctx context.Context, key string, member string) (float64, error) {
	return r.Redis.ZScore(ctx, key, member).Result()
}

func (r *RedisClient) ZAdd(ctx context.Context, key string, score float64, member string) error {
	return r.Redis.ZAdd(ctx, key, &redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

// LPush List
func (r *RedisClient) LPush(ctx context.Context, key string, values string) error {
	return r.Redis.LPush(ctx, key, values).Err()
}

func (r *RedisClient) LRem(ctx context.Context, key string, count int64, value interface{}) error {
	return r.Redis.LRem(ctx, key, count, value).Err()
}

func (r *RedisClient) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.Redis.LRange(ctx, key, start, stop).Result()
}

func (r *RedisClient) SetNX(ctx context.Context, key string, value interface{}, expire time.Duration) (bool, error) {
	return r.Redis.SetNX(ctx, key, value, expire).Result()
}

func (r *RedisClient) Expire(ctx context.Context, key string, expire time.Duration) (bool, error) {
	return r.Redis.Expire(ctx, key, expire).Result()
}
