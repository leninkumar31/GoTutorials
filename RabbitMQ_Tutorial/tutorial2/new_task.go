package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// make a connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Not able to connect to RabbitMQ")
	defer conn.Close()
	// Create a channel
	ch, err := conn.Channel()
	handleError(err, "Not able to create a channel")
	defer ch.Close()
	// Declare a queue
	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Not able to Declare queue")
	// Publish the task
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	handleError(err, "Failed to publish")
	log.Printf(" [x] sent %s:", body)
}

func bodyFrom(args []string) string {
	var s string
	if len(args) < 2 || args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
