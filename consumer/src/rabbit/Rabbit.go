package rabbit

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
    Broker  *amqp.Connection
    Channel *amqp.Channel
}

// conn rabbit
func NewRabbit() *Rabbit {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }
	 
   
    rabbitUrl := os.Getenv("RABBIT_URL")

    
    conn, err := amqp.Dial(rabbitUrl)
    if err != nil {
        log.Fatal("Error al abrir una conexión hacia RabbitMQ")
    }

    // abre un canal
    ch, err := conn.Channel()
    if err != nil {
        log.Fatal("Error al abrir un canal")
    }

    return &Rabbit{Broker: conn, Channel: ch}
}

// declara el exchange y la cola
func (r *Rabbit) SetupExchangeAndQueues() {
    // exchange
    err := r.Channel.ExchangeDeclare(
        "inscriptions_exchange", // name del exchange
        "direct",               // Tipo de exchange
        true,                   // Durable
        false,                  // Auto-deleted
        false,                  // Internal
        false,                  // No-wait
        nil,                    // Arguments
    )
    FailOnError(err, "Error al declarar el exchange")

    //  la cola
    _, err = r.Channel.QueueDeclare(
        "inscriptionsQueue", // name de la cola
        true,               // Durable
        false,              // Delete when unused
        false,              // Exclusive
        false,              // No-wait
        nil,                // Arguments
    )
    FailOnError(err, "Error al declarar la cola de inscripciones")

    // vincula la cola con el exchange
    err = r.Channel.QueueBind(
        "inscriptionsQueue",         // name de la cola
        "inscriptions_routing_key", // Routing key
        "inscriptions_exchange",    // name del exchange
        false,                      // No-wait
        nil,                        // Arguments
    )
    FailOnError(err, "Error al vincular la cola de inscripciones")
}

//  consume los mensajes de la cola 
func (r *Rabbit) ReceiveContent() <-chan amqp.Delivery {
    msgs, err := r.Channel.Consume(
        "inscriptionsQueue", // name de la cola
        "",                 // Consumer
        true,               // Auto-ack
        false,              // Exclusive
        false,              // No-local
        false,              // No-wait
        nil,                // Args
    )
    FailOnError(err, "Error al registrar el consumidor")

    return msgs
}

func FailOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}