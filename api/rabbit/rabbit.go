package rabbit

import (
	"log"
	"fmt"
	"strings"
	"github.com/streadway/amqp"
	"../../services/matching"
)

var (
	sendCh *amqp.Channel
	receiveCh *amqp.Channel
	receiveQ amqp.Queue
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {

	sendConn, err := amqp.Dial("amqp://cduica:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer sendConn.Close()

	sendCh, err = sendConn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer sendCh.Close()

	err = sendCh.ExchangeDeclare(
		"cduica-hello",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	receiveConn, err := amqp.Dial("amqp://cduica:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer receiveConn.Close()

	receiveCh, err = receiveConn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer receiveCh.Close()

	err = receiveCh.ExchangeDeclare(
		"cduica-world",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	receiveQ, err = receiveCh.QueueDeclare(
		"", // name
		false,   // durable
		true,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func PublishMatch( userA string, userB string ) {
	body := userA + " " + userB
	err := sendCh.Publish(
		"cduica-hello",     // exchange
		"", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf("[x] Sent message %s", body)
	failOnError(err, "Failed to publish a message")
}

func Consume() {

	er := receiveCh.QueueBind(
		receiveQ.Name, // queue name
		"",     // routing key
		"cduica-world", // exchange
		false,
		nil,
	)
	failOnError(er, "Failed to bind a queue")

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
		  log.Printf( "[x] Received a message: %s", d.Body )
		  s := string( d.Body[:] )
		  split := strings.Split( s, " " )
		  if len(split) < 2 {
			  //panic("message incorrect")
			  continue
		  }
		  err, match := matching.SaveMatch( split[0], split[1] )
		  if err != nil {
			  fmt.Println(err)
			  continue
		  }
		  if match.Matched == true {
			  PublishMatch(split[0], split[1])
			  PublishMatch(split[1], split[0])
		  }
		}
	  }()
	  
	  log.Println("RabbitMQ started \n")
	  <-forever
}