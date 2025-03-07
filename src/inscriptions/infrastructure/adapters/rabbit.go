package adapters

import (
	"github.com/Andresito126/go-validation_inscriptions/src/core/rabbitmq"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/repository"
	"github.com/streadway/amqp"
)

type RabbitAdapter struct {
    conn *amqp.Connection
    ch   *amqp.Channel
}

func NewRabbitAdapter() (repository.IRabbitService, error) {
    conn, err := rabbitmq.ConnectRabbitMQ()
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    return &RabbitAdapter{
        conn: conn,
        ch:   ch,
    }, nil
}

func (r *RabbitAdapter) Consume(queue string, handler func([]byte)) error {
    q, err := r.ch.QueueDeclare(
        "inscriptionsQueue", // name
        true, // durable
        false, // delete when unused
        false, // exclusive
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        return err
    }

    msgs, err := r.ch.Consume(
        q.Name, // name
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        return err
    }

    for msg := range msgs {
        handler(msg.Body) // Procesar el mensaje
    }

    return nil
}