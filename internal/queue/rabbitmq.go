package queue

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQ() (*amqp091.Connection, *amqp091.Channel) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	_, err = ch.QueueDeclare(
		"listing_created",
		true,
		false, 
		false, 
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ch.QueueDeclare(
		"listing_failed",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	return conn, ch 
}