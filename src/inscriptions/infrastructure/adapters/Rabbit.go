package adapters

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Broker  *amqp.Connection
	Channel *amqp.Channel
}

// NewRabbit establece la conexión a RabbitMQ
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

	// Abrir un canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error al abrir un canal")
	}

	return &Rabbit{Broker: conn, Channel: ch}
}

// declara el exchange, la cola y su vinculo para las notificaciones
func (r *Rabbit) SetupNotificationExchangeAndQueue() {
	// Declarar el exchange "notificationsExchange"
	err := r.Channel.ExchangeDeclare(
		"notificationsExchange", // Name 
		"direct",                // Type de exchange
		true,                    // Durable
		false,                   // Auto-deleted
		false,                   // Internal
		false,                   // No-wait
		nil,                     // Arguments
	)
	FailOnError(err, "Error al declarar el exchange de notificaciones")

	// Declarar la cola "notificationsQueue"
	_, err = r.Channel.QueueDeclare(
		"notificationsQueue", // Name de la cola 
		true,                 // Durable
		false,                // Delete when unused
		false,                // Exclusive
		false,                // No-wait
		nil,                  // Arguments
	)
	FailOnError(err, "Error al declarar la cola de notificaciones")

	// Vincular la cola con el exchange usando la routing key 
	err = r.Channel.QueueBind(
		"notificationsQueue",  // Name de la cola 
		"notification",        // Routing key 
		"notificationsExchange", // Name 
		false,                 // No-wait
		nil,                   // Arguments
	)
	FailOnError(err, "Error al vincular la cola de notificaciones con el exchange")
}

//  publica el mensaje en la cola
func (r *Rabbit) SendMessageToBroker(studentID, courseID, status string) {
	// crea el mensaje

	var data = map[string]string{
		"student_id": studentID,
		"course_id": courseID,
		"status": status,
	}

	m, err := json.Marshal(data)

	if err != nil {
		log.Fatalf("Error al serializar el mensaje: %v", err)
	}

	// message := "StudentID: " + studentID + " CourseID: " + courseID + " Status: " + status

	// publicar el mensaje en la cola 
	err = r.Channel.Publish(
		"notificationsExchange", // Exchange 
		"notification",          // Routing key 
		false,                   // Mandatory
		false,                   // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(m),
		},
	)

	// verificar si hubo un error al enviar el mensaje
	if err != nil {
		log.Printf("Error al enviar mensaje a la cola: %s", err)
	} else {
		log.Printf("Mensaje enviado a la cola correctamente: %s", data)
	}
}

// FailOnError maneja los errores
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
