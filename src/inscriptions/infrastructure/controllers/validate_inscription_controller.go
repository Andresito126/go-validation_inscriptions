// src/api2/infrastructure/controllers/validate_inscription_controller.go
package controllers

import (
    "encoding/json"
    "log"

    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/usecases"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/domain/entities"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/services"
)

type ValidateInscriptionController struct {
    useCase        *usecases.ValidateInscriptionUseCase
    rabbitService  *services.RabbitService
}

func NewValidateInscriptionController(useCase *usecases.ValidateInscriptionUseCase, rabbitService *services.RabbitService) *ValidateInscriptionController {
    return &ValidateInscriptionController{
        useCase:       useCase,
        rabbitService: rabbitService,
    }
}

func (c *ValidateInscriptionController) Start() {
    log.Println("Iniciando el consumo de mensajes de RabbitMQ...")
    err := c.rabbitService.Consume("inscriptionsQueue", c.Handle)
    if err != nil {
        log.Fatalf("Error al consumir mensajes: %v", err)
    }
}

func (c *ValidateInscriptionController) Handle(message []byte) {
    log.Printf("Mensaje recibido: %s", string(message))

    var inscription entities.Inscription
    if err := json.Unmarshal(message, &inscription); err != nil {
        log.Printf("Error al deserializar el mensaje: %v", err)
        return
    }

    if inscription.ID <= 0 {
        log.Printf("ID de inscripción inválido: %d", inscription.ID)
        return
    }

    log.Printf("Procesando inscripción: ID=%d, StudentID=%d, CourseID=%d", inscription.ID, inscription.StudentID, inscription.CourseID)

    // valida la inscripción 
    if err := c.useCase.Run(&inscription); err != nil {
        log.Printf("Error al validar la inscripción: %v", err)
        return
    }

    
    log.Printf("Inscripción validada: ID=%d, Status=%s", inscription.ID, inscription.Status)
}