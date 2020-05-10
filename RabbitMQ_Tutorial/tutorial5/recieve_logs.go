package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s; %s", err, msg)
	}
}

func main() {
	// Establish connection with RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	handleError(err, "Unable to establish connection with rabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	handleError(err, "Unable to create channel")
	defer ch.Close()

	// Declare exchange
	err = ch.ExchangeDeclare(
		"logs_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Unable to declare an exchange")

	// Declare queue
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	handleError(err, "unable to declare queue")

	// Bind the channel with queue
	args := os.Args
	if len(args) < 2 || args[1] == "" {
		log.Printf("Usage: Binding Key....%s", args[0])
		os.Exit(0)
	}
	for _, s := range args[1:] {
		log.Printf("binding exchange with queue for binding key: %s", s)
		err = ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
		)
		handleError(err, "unable to bind queue and exchange")
	}
	// Register consumer with queue
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Unable to register consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Message recieved: %s", d.Body)
		}
	}()
	log.Printf("Waiting for recieveing")
	<-forever
}
