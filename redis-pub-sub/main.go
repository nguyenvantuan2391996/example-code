package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func initRedis() (*RedisClient, error) {
	var redisClient *redis.Client

	opts, err := redis.ParseURL("redis://default:@localhost:6379")
	if err != nil {
		log.Fatal("failed to init redis:", err)
		return nil, err
	}

	opts.PoolSize = 30
	opts.DialTimeout = 10 * time.Second
	opts.ReadTimeout = 5 * time.Second
	opts.WriteTimeout = 5 * time.Second
	opts.IdleTimeout = 5 * time.Second
	opts.Username = ""

	redisClient = redis.NewClient(opts)

	cmd := redisClient.Ping(context.Background())
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	return &RedisClient{
		Redis: redisClient,
	}, nil
}

func main() {
	// init redis
	redisClient, err := initRedis()
	if err != nil {
		logrus.Warnf("init redis client is failed with err: %v", err)
		return
	}

	ctx, channel := context.Background(), "test"
	ch := make(chan string, 1)
	numberOfWorkers := 5
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 5 workers
	for i := 0; i < numberOfWorkers; i++ {
		go func() {
			for msg := range ch {
				fmt.Println(msg)
			}
		}()
	}

	subscriber := redisClient.Subscribe(ctx, channel)
	go func() {
		for {
			message, err := subscriber.ReceiveMessage(ctx)
			if err != nil {
				continue
			}

			// push to channel
			ch <- message.Payload
		}
	}()

	time.Sleep(1 * time.Second)
	go func() {
		for i := 0; i < 1000; i++ {
			redisClient.Redis.Publish(ctx, channel, fmt.Sprintf("hello %v", i))
		}
	}()

	<-quit
	log.Println("shutting down")
}
