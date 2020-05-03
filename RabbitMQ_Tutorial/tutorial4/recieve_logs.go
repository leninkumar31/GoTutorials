package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	handleError(err, "Unable to establish connection")
	defer conn.Close()

	ch, err := conn.Channel()
	handleError(err, "Unable to create channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Unable to declare a channel")

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	handleError(err, "Unable to declare Queue")
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("You haven't provided severity type")
	}
	for _, s := range args {
		log.Printf("Binding queue % with exchange %s with routing key %s", q.Name, "logs_direct", s)
		err = ch.QueueBind(
			q.Name,
			s,
			"logs_direct",
			false,
			nil,
		)
		handleError(err, "Fialed to bind queue")
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Unable to register consumer with the queue")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Message recieved: %s", d.Body)
		}
	}()
	log.Printf("Waiting for message")
	<-forever
}
