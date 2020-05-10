package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func main() {
	// Establish connection with RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	handleError(err, "Unable to establish connection with RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	handleError(err, "Unable to create a Channel")
	defer ch.Close()

	// Declare Exchange
	err = ch.ExchangeDeclare(
		"logs_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Unable to declare exchange")
	args := os.Args
	body := getBodyFrom(args)
	err = ch.Publish(
		"logs_topic",
		severityFrom(args),
		false,
		false,
		amqp.Publishing{
			ContentType: "plain/text",
			Body:        []byte(body),
		})
	handleError(err, "Message not Published")
	log.Printf("Message sent: %s", body)
}

func getBodyFrom(args []string) string {
	var s string
	if len(args) < 3 || args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if len(args) < 2 || args[1] == "" {
		s = "anonymous.info"
	} else {
		s = args[1]
	}
	return s
}
