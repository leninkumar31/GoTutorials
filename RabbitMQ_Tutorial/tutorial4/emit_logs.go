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
	// Establish connection with RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	handleError(err, "Unable to establish Connection")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	handleError(err, "Unable to create a channel")
	defer ch.Close()

	// Declare an Exchange
	err = ch.ExchangeDeclare(
		"logs_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Unable to declare an Exchange")

	// Publish message
	args := os.Args
	body := bodyFrom(args)
	err = ch.Publish(
		"logs_direct",
		severityFrom(args),
		false,
		false,
		amqp.Publishing{
			ContentType: "plain/text",
			Body:        []byte(body),
		})
	handleError(err, "Unable to publish message")
	log.Printf("Message sent: %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if len(args) < 3 || args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

// info, warning or error
func severityFrom(args []string) string {
	var s string
	if len(args) < 2 || args[1] == "" {
		s = "info"
	} else {
		s = args[1]
	}
	return s
}
