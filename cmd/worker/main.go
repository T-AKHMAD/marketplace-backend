package main

import (
	"encoding/json"
	"log"
	"marketplace/internal/domain"
	"marketplace/internal/queue"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, ch := queue.NewRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		"listing_created",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for msg := range msgs {
		var l domain.Listing

		if err := json.Unmarshal(msg.Body, &l); err != nil {
			log.Printf("failed to unmarshal message: %v", err)

			err = ch.Publish(
				"",
				"listing_failed",
				false,
				false,
				amqp091.Publishing{
					ContentType: "application/json",
					Body:        msg.Body,
				},
			)
			if err != nil {
				log.Printf("failed to publish to DLQ: %v", err)
				continue
			}
			msg.Ack(false)
			continue
		}

		log.Printf("Listing created: ID=%d Title=%s Price=%d", l.ID, l.Title, l.Price)
		msg.Ack(false)
	}
}
