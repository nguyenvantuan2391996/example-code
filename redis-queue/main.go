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

	ctx, queue := context.Background(), "tuan"
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// the worker consumes queue
	go func() {
		for {
			// using BRPop will wait with a timeout if the queue is empty. If timeout is 0 it will wait forever
			message, err := redisClient.BRPop(context.Background(), queue, 0)
			if err != nil {
				fmt.Println(err)
				continue
			}

			time.Sleep(5 * time.Second)
			fmt.Println(fmt.Sprintf("message: %v", message))
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			_, err := redisClient.LPush(ctx, queue, fmt.Sprintf("hello %v", i))
			if err != nil {
				return
			}
		}
	}()

	<-quit
	log.Println("shutting down")
}
