package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnFatalError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func main() {
	// connecting to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnFatalError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// create a channel
	ch, err := conn.Channel()
	failOnFatalError(err, "Failed to open a channel")
	defer ch.Close()
	// Declare a queue to publish messages
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnFatalError(err, "failed to create a queue")
	// Publish message to the queue
	body := "Hello Subscriber!!"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf("[x] sent %s", body)
	failOnFatalError(err, "Failed to publish a message")
}
