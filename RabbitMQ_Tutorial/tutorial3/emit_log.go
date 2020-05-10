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
	//Establish connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	handleError(err, "Unable to connet to RabbitMQ")
	defer conn.Close()

	//Create a channel
	ch, err := conn.Channel()
	handleError(err, "Unable to create a channel")
	defer ch.Close()

	//Declare an Exchange
	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Unable to declare an exchange")
	args := os.Args
	body := getBody(args)

	//Publish the logs
	err = ch.Publish(
		"logs",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "plain/text",
			Body:        []byte(body),
		})
	handleError(err, "unable to publish message")
	log.Printf("Message published: %s", body)
}

func getBody(args []string) string {
	var s string
	if len(args) < 2 || args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
