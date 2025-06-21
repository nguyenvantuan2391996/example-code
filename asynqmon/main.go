package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "0.0.0.0:6379", Password: ""})
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	payload := []byte("Hello World")

	info, err := client.Enqueue(asynq.NewTask("hihi", payload), asynq.ProcessIn(10*time.Second))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	go func() {
		srv := asynq.NewServer(
			asynq.RedisClientOpt{Addr: "0.0.0.0:6379", Password: ""},
			asynq.Config{
				// Specify how many concurrent workers to use
				Concurrency: 10,
				// Optionally specify multiple queues with different priority.
				Queues: map[string]int{
					"critical": 6,
					"default":  3,
					"low":      1,
				},
				// See the godoc for other configuration options
				TaskCheckInterval: 5 * time.Second,
			},
		)

		mux := asynq.NewServeMux()
		mux.HandleFunc("hihi", func(ctx context.Context, task *asynq.Task) error {
			fmt.Printf("hihi: %s", task.Payload())
			return nil
		})

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	}()

	time.Sleep(time.Hour)
}
