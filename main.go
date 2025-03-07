package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	
	conn, err := amqp.Dial("amqp://andreju:andreju@44.220.150.133:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"inscriptionQueue", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")


	//aca seria el proceso de recibir la cola, o sea consumirla


	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	
	var forever chan struct{}
	
	go func() {
		for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		}
	}()
	
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	//esta es la forma de mantener el programa corriendo con el forever channel

	<-forever

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
