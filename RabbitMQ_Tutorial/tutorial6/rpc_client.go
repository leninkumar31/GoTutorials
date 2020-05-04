package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func fibanocciRPC(n int) (res int, err error) {
	// Establish connection with RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	handleError(err, "Unable to establish connection with RabbitMQ")
	defer conn.Close()
	// Create a channel
	ch, err := conn.Channel()
	handleError(err, "Unable to Create a channel")
	defer ch.Close()
	// Declare callback queue
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	handleError(err, "Unable to declare queue")
	// Register consumer for RPC response
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "failed to register consumer")
	// Publish the RPC request
	corrID := randomString(32)
	err = ch.Publish(
		"",
		"rpc_queue",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	handleError(err, "Unable to publish the request")

	for d := range msgs {
		if corrID == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			handleError(err, "failed to convert Body to String")
			break
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	args := os.Args
	n := bodyFrom(args)
	log.Printf(" [x] Requesting Fib(%d)", n)
	res, err := fibanocciRPC(n)
	handleError(err, "Failed to handle RPC request")
	log.Printf(" [.] Got %d", res)
}

func bodyFrom(args []string) int {
	var s string
	if len(args) < 2 || args[1] == "" {
		s = "30"
	} else {
		s = args[1]
	}
	n, err := strconv.Atoi(s)
	handleError(err, "Failed to conv arg[1] to integer")
	return n
}
