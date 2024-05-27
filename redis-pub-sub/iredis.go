package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -package=main -destination=iredis_mock.go -source=iredis.go
type IRedisClientInterface interface {
	Subscribe(ctx context.Context, channel string) *redis.PubSub
	Publish(ctx context.Context, channel, message string) (int64, error)
}
