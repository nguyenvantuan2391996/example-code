package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	Master = "master"
	Slave  = "slave"
)

type Config struct {
	MasterURL string
	SlavesURL []string
}

type RedisResolver struct {
	Master *redis.Client
	Salves []*redis.Client
}

func open(ctx context.Context, dsn string) (*redis.Client, error) {
	var redisClient *redis.Client

	opts, err := redis.ParseURL(dsn)
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

	cmd := redisClient.Ping(ctx)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	return redisClient, nil
}

func NewRedisResolver(config *Config) (RedisResolverInterface, error) {
	// Open redis master
	master, err := open(context.Background(), config.MasterURL)
	if err != nil {
		return nil, err
	}

	// Open redis slaves
	slaves := make([]*redis.Client, 0)
	if len(config.SlavesURL) == 0 {
		slaves = append(slaves, master)
	} else {
		for _, dsn := range config.SlavesURL {
			slave, errSlave := open(context.Background(), dsn)
			if errSlave != nil {
				return nil, errSlave
			}

			slaves = append(slaves, slave)
		}
	}

	return &RedisResolver{
		Master: master,
		Salves: slaves,
	}, nil
}

func (rs *RedisResolver) Clauses(action string) IRedisClientInterface {
	client := &redis.Client{}

	switch action {
	case Master:
		client = rs.Master
	case Slave:
		client = rs.Salves[rand.Intn(len(rs.Salves))]
	}

	return &RedisClient{Redis: client}
}
