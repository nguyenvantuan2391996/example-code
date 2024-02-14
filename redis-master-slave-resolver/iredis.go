package main

import (
	"context"
	"time"
)

//go:generate mockgen -package=main -destination=iredis_mock.go -source=iredis.go
type IRedisClientInterface interface {
	Set(ctx context.Context, key string, value interface{}, expire time.Duration) (string, error)
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) (int64, error)
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	Exists(ctx context.Context, keys ...string) (int64, error)
	ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error)
	ZScore(ctx context.Context, key string, member string) (float64, error)
	ZAdd(ctx context.Context, key string, score float64, member string) error
	LPush(ctx context.Context, key string, values string) error
	LRem(ctx context.Context, key string, count int64, value interface{}) error
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	SetNX(ctx context.Context, key string, value interface{}, expire time.Duration) (bool, error)
	Expire(ctx context.Context, key string, expire time.Duration) (bool, error)
}
