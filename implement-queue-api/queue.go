package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	NumberOfWorkers  = 10
	NumberOfCapacity = 10
)

type Queue struct {
	Channel   chan string
	WaitGroup sync.WaitGroup
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		Channel: make(chan string, capacity),
	}
}

func (q *Queue) Enqueue(request string) {
	q.Channel <- request
}

func (q *Queue) Dequeue() {
	for i := 0; i < NumberOfWorkers; i++ {
		q.WaitGroup.Add(1)
		go func() {
			defer q.WaitGroup.Done()
			for msg := range q.Channel {
				time.Sleep(5 * time.Second)
				fmt.Printf("handle message %v", msg)
			}
		}()
	}
}
