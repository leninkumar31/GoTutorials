package main

import (
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func fibanacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibanacci(n-1) + fibanacci(n-2)
}

func main() {
	// Establish connection with RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	handleError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	handleError(err, "Failed to create a channel")
	defer ch.Close()

	// Declare queue
	q, err := ch.QueueDeclare(
		"rpc_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Failed to declare queue")

	err = ch.Qos(
		1,
		0,
		false,
	)
	handleError(err, "Failed to create Qos")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			handleError(err, "Failed to convert body to integer")
			log.Printf("Fibanacci of %d", n)
			res := fibanacci(n)
			err = ch.Publish(
				"",
				d.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(res)),
				})
			handleError(err, "failed to publish")
			d.Ack(false)
		}
	}()
	log.Printf("Waiting for RPC requests")
	<-forever
}
