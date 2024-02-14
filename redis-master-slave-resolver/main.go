package main

import (
	"context"
	"fmt"
)

func main() {
	master := "redis://default:@localhost:6279"
	slaves := []string{
		"redis://default:@localhost:6179",
		"redis://default:@localhost:6079",
	}

	redisResolver, err := NewRedisResolver(&Config{
		MasterURL: master,
		SlavesURL: slaves,
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	key := "test"
	_, err = redisResolver.Clauses(Master).Set(ctx, key, "hello world!", 0)
	if err != nil {
		panic(err)
	}

	// read from master
	m, err := redisResolver.Clauses(Master).Get(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("from master %v", m))

	// read from slave
	s, err := redisResolver.Clauses(Slave).Get(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("from slave %v", s))

	// error slave write
	_, err = redisResolver.Clauses(Slave).Set(ctx, key, "hello", 10)
	fmt.Println(err)
}
