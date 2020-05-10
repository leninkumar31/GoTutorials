package main

import (
	"log"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Establish connection with RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	handleError(err, "Unable to Establish connection with RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	handleError(err, "unable to create a channel")
	defer ch.Close()

	// Decalre Exchange
	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "unable to declare an exchange")

	// Declare Queue
	q, err := ch.QueueDeclare(
		"",
		true,
		false,
		true,
		false,
		nil,
	)
	handleError(err, "Unable to declare queue")

	// Bind exchnage with Queue
	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	handleError(err, "unable to bind exchange and queue")

	// Consume messages
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
			log.Printf("Reieved Message :%s", d.Body)
		}
	}()
	log.Printf("Waiting for logs. To exit press CTRL + C")
	<-forever
}
