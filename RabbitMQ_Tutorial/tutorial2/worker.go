package main

import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Establish connection with RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "unable to connect to RabbitMQ")
	defer conn.Close()
	// Make a channel
	ch, err := conn.Channel()
	handleError(err, "Unable to create a channel")
	defer ch.Close()
	// Feclare queue
	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Not able to declare queue")
	// Set Prefetch Count
	err = ch.Qos(
		1,
		0,
		false,
	)
	// consume messages
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
			log.Printf("Message recieved: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()
	log.Printf("Waiting for message")
	<-forever
}
