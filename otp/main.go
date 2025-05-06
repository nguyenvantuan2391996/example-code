package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type OTP struct {
	RedisClient *redis.Client
	ExpiredTime time.Duration
	OTPLength   int
}

func NewOTP(redisClient *redis.Client) *OTP {
	return &OTP{
		RedisClient: redisClient,
		ExpiredTime: 5 * time.Second,
		OTPLength:   6,
	}
}

func initRedis() (*redis.Client, error) {
	var redisClient *redis.Client

	opts, err := redis.ParseURL("redis://default:oK9T4D3xOfWcIOB@0.0.0.0:6377")
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

func (o *OTP) GenerateOTP(length int) (string, error) {
	const otpChars = "1234567890"

	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func (o *OTP) HandleOTP(userCode string) (string, error) {
	otp, err := o.GenerateOTP(o.OTPLength)
	if err != nil {
		return "", err
	}

	err = o.RedisClient.Set(context.Background(), fmt.Sprintf("otp-%v", userCode), otp, o.ExpiredTime).Err()
	if err != nil {
		return "", err
	}

	return otp, nil
}

func (o *OTP) VerifyOTP(otpRequest, userCode string) error {
	otp, err := o.RedisClient.Get(context.Background(), fmt.Sprintf("otp-%v", userCode)).Result()
	if err != nil {
		return err
	}

	if otp != otpRequest {
		return fmt.Errorf("invalid otp request")
	}

	return nil
}

func main() {
	client, err := initRedis()
	if err != nil {
		log.Fatal(err)
	}

	otpService := NewOTP(client)

	otp, err := otpService.HandleOTP("123456")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			errVerify := otpService.VerifyOTP(otp, "123456")
			switch {
			case errVerify == nil:
				log.Println("the otp is valid")
			case errors.Is(errVerify, redis.Nil):
				log.Println("the otp is expired")
			case errVerify.Error() == "invalid otp request":
				log.Println("the otp is invalid")
			default:
				log.Fatal("system error:", errVerify)
			}

			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Hour)
}
