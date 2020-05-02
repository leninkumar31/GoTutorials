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
	// Declare a queue to recieve messages
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnFatalError(err, "failed to create a queue")
	// Consume message to the queue
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnFatalError(err, "failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Recieved a message: %s", d.Body)
		}
	}()
	log.Printf("Waiting for the message")
	<-forever
}
