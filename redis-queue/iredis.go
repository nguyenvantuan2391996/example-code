package main

import (
	"context"
	"time"
)

//go:generate mockgen -package=main -destination=iredis_mock.go -source=iredis.go
type IRedisClientInterface interface {
	LPush(ctx context.Context, queue string, value interface{}) (int64, error)
	BRPop(ctx context.Context, queue string, timeout time.Duration) ([]string, error)
}
