package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Andresito126/ConsumerRabbit/src/models"
	"github.com/Andresito126/ConsumerRabbit/src/rabbit"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	
	rabbit := rabbit.NewRabbit()

	// obtiene los mensajes de rabbit
	msgs := rabbit.ReceiveContent()

	// Procesar los mensajes
	ProcessMessage(msgs)
}

// par procesar los mensajes de rabbit
func ProcessMessage(msgs <-chan amqp.Delivery) {
	forever := make(chan struct{})

	//  el goroutine para pdoer procesar los mensajes
	go func() {
		for d := range msgs {
			var inscription models.Inscription

			
			err := json.Unmarshal(d.Body, &inscription)
			if err != nil {
				log.Printf("Error al decodificar el mensaje: %s", err)
				continue
			}

			log.Printf("[x] Inscripción recibida: ID Estudiante %d, ID Curso %d", inscription.StudentID, inscription.CourseID)

			// realiza la peticion para la validacion ala api2
			Fetch(inscription)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// realizar el post a la api2
func Fetch(inscription models.Inscription) {
	url := "http://localhost:8081/inscriptions/validate" 

	// vualve el objeto a JSON
	jsonPayload, err := json.Marshal(inscription)
	if err != nil {
		log.Fatalf("Error al serializar la inscripción: %v", err)
	}
	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("Error al hacer la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Error al procesar la inscripción, código: %d", resp.StatusCode)
	}

	//lee lo que viene de la api
	var result struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Error al decodificar la respuesta de la API: %v", err)
	}

	// respuesta hace::
	if result.Status == "aceptada" {
		log.Println("Inscripción aceptada")
	} else {
		log.Println("Inscripción rechazada")
	}
}
