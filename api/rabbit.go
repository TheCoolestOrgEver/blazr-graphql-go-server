package rabbit

import (
	"log"
	"strings"
	"github.com/streadway/amqp"
	"../services/matching"
)

var (
	sendCh amqp.Channel
	sendQ amqp.Queue
	receiveCH amqp.Channel
	receiveQ amqp.Queue
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {

	sendConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer sendConn.Close()

	sendCh, err = sendConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer sendCh.Close()

	sendQ, err = sendCh.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	receiveConn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer receiveConn.Close()

	receiveCh, err := receiveConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer receiveCh.Close()

	receiveQ, err := receiveCh.QueueDeclare(
	"hello", // name
	false,   // durable
	false,   // delete when usused
	false,   // exclusive
	false,   // no-wait
	nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func PublishMatch( userA string, userB string ) {
	body := userA + " " + userB
	err = sendCh.Publish(
		"",     // exchange
		sendQ.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func Consume() {
	msgs, err := receiveCh.Consume(
		receiveQ.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	  )
	  failOnError(err, "Failed to register a consumer")
	  
	  forever := make(chan bool)
	  
	  go func() {
		for d := range msgs {
		  log.Printf( "Received a message: %s", d.Body )
		  s := string( d.Body[:] )
		  split := strings.Split( s, " " )
		  match := matching.SaveMatch( split[0], split[1] )
		  if match.Matched {
			  PublishMatch(split[1], split[0])
		  }
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever
}