package controllers

import (
	"net/http"
	"strconv"

	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/services"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/usecases"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/domain/entities"
	"github.com/gin-gonic/gin"
)

type ValidateInscriptionController struct {
    validateUseCase    *usecases.ValidateInscriptionUseCase
    updateStatusUseCase *usecases.UpdateInscriptionStatusUseCase
    eventService       *services.EventService
}

func NewValidateInscriptionController(
    validateUseCase *usecases.ValidateInscriptionUseCase,
    updateStatusUseCase *usecases.UpdateInscriptionStatusUseCase,
    eventService *services.EventService,
) *ValidateInscriptionController {
    return &ValidateInscriptionController{
        validateUseCase:    validateUseCase,
        updateStatusUseCase: updateStatusUseCase,
        eventService:       eventService,
    }
}

func (c *ValidateInscriptionController) HandleHTTP(ctx *gin.Context) {
    var inscription entities.Inscription

    if err := ctx.ShouldBindJSON(&inscription); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    // validacion de la inscripción
    status, err := c.validateUseCase.Run(&inscription)
    if err != nil {
        if status == "rechazada" {
            ctx.JSON(http.StatusOK, gin.H{"status": "rechazada", "message": err.Error()})
        } else {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating inscription"})
        }
        return
    }

    // actu en la base de datos
    if err := c.updateStatusUseCase.Run(&inscription, status); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating inscription status"})
        return
    }

    // punlciador de notificación, se  int a string
    c.eventService.PublishNotification(
        strconv.Itoa(inscription.StudentID), 
        strconv.Itoa(inscription.CourseID),  
        status,
    )

    // Respuesta
    ctx.JSON(http.StatusOK, gin.H{"status": status})
}
