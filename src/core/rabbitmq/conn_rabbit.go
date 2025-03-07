package rabbitmq

import (
    "log"
    "os"
    "sync"

    "github.com/joho/godotenv"
    "github.com/streadway/amqp"
)

var (
    rabbitInstance *amqp.Connection
    rabbitOnce     sync.Once
)

func ConnectRabbitMQ() (*amqp.Connection, error) {
    rabbitOnce.Do(func() {
        if err := godotenv.Load(); err != nil {
            log.Fatalf("Error al cargar el archivo .env: %v", err)
        }

        user := os.Getenv("RABBITMQ_USER")
        password := os.Getenv("RABBITMQ_PASSWORD")
        host := os.Getenv("RABBITMQ_HOST")
        port := os.Getenv("RABBITMQ_PORT")

        conn, err := amqp.Dial("amqp://" + user + ":" + password + "@" + host + ":" + port + "/")
        if err != nil {
            log.Fatalf("Error al conectar a RabbitMQ: %v", err)
        }

        rabbitInstance = conn
        log.Println("Conexión a RabbitMQ exitosa")
    })

    return rabbitInstance, nil
}