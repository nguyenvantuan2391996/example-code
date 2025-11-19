package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	Ch *amqp091.Channel
}

func NewRabbitMQ() (*RabbitMQClient, error) {
	conn, err := amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return &RabbitMQClient{
		Ch: ch,
	}, nil
}

func (r *RabbitMQClient) Consumer(queue string) {
	msgs, err := r.Ch.Consume(
		queue,
		"",
		false, // auto-ack
		false, // exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	go func() {
		for d := range msgs {
			time.Sleep(5 * time.Second)
			log.Printf("Received: %s", d.Body)
			err = d.Ack(false)
			if err != nil {
				return
			}
		}
	}()

	log.Println("Waiting for messages...")
}

func (r *RabbitMQClient) Publish(queue, exchange, body string) {
	q, err := r.Ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = r.Ch.PublishWithContext(ctx,
		exchange, // exchange
		q.Name,   // routing key
		false,    // mandatory
		false,    // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	log.Printf("Sent: %s", body)
}

func main() {
	rb, err := NewRabbitMQ()
	if err != nil {
		panic(err)
	}

	rb.Consumer("test")

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			rb.Publish("test", "", fmt.Sprintf("%v", t))
		}
	}
}
