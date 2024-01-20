package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func initRedis() (*redis.Client, error) {
	var redisClient *redis.Client

	opts, err := redis.ParseURL("redis://default:admin123@localhost:6379")
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

	return redisClient, nil
}

func DoNotUseLuaScript(redisClient *redis.Client) int {
	// Step 1
	currentCountStr, err := redisClient.Get(context.Background(), "counter").Result()
	if err != nil && err != redis.Nil {
		logrus.Warnf("Get value redis by key %v is failed with err: %v", "counter", err)
		return 0
	}

	if err == redis.Nil {
		currentCountStr = "0"
	}

	// Step 2
	currentCount, err := strconv.Atoi(currentCountStr)
	if err != nil {
		logrus.Warnf("Convert string to int is failed with err: %v", err)
		return 0
	}

	// Step 3: increase
	newCount := currentCount + 1

	// Step 4: Set value redis
	isSet := redisClient.Set(context.Background(), "counter", newCount, 0).Val()
	if isSet != "OK" {
		logrus.Warnf("Set value redis by key %v is failed with err: %v", "counter", err)
		return 0
	}

	return newCount
}

func UseLuaScript(redisClient *redis.Client) int {
	script := `
		local currentCount = tonumber(redis.call('GET', KEYS[1]) or '0')
		local newCount = currentCount + tonumber(ARGV[1])
		redis.call('SET', KEYS[1], newCount)
		return newCount
	`

	// Running Lua Script
	resultStr, err := redisClient.Eval(context.Background(), script, []string{"counter-lua"}, 1).Result()
	if err != nil {
		logrus.Warnf("Running Lua Script is failed with err: %v", err)
		return 0
	}

	result, err := strconv.Atoi(fmt.Sprintf("%v", resultStr))
	if err != nil {
		logrus.Warnf("Convert string to int is failed with err: %v", err)
		return 0
	}

	return result
}

func main() {
	// init redis
	redisClient, err := initRedis()
	if err != nil {
		logrus.Warnf("init redis client is failed with err: %v", err)
		return
	}

	// use lua script
	resultLua := 0
	startLua := time.Now()
	for i := 0; i < 100; i++ {
		resultLua = UseLuaScript(redisClient)
	}
	endLua := time.Now()
	logrus.Infof("Result lua script type: %v and time %v ms", resultLua, endLua.Sub(startLua).Milliseconds())

	// do not use lua script
	result := 0
	start := time.Now()
	for i := 0; i < 100; i++ {
		result = DoNotUseLuaScript(redisClient)
	}
	end := time.Now()
	logrus.Infof("Result do not use lua script type: %v and time %v ms", result, end.Sub(start).Milliseconds())
}
