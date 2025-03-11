package controllers

import (
    
    "net/http"

    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/usecases"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/domain/entities"
    "github.com/gin-gonic/gin"
)

type ValidateInscriptionController struct {
    validateUseCase    *usecases.ValidateInscriptionUseCase
    updateStatusUseCase *usecases.UpdateInscriptionStatusUseCase
}

func NewValidateInscriptionController(validateUseCase *usecases.ValidateInscriptionUseCase, updateStatusUseCase *usecases.UpdateInscriptionStatusUseCase) *ValidateInscriptionController {
    return &ValidateInscriptionController{
        validateUseCase:    validateUseCase,
        updateStatusUseCase: updateStatusUseCase,
    }
}

// handle para los casos de uso
func (c *ValidateInscriptionController) HandleHTTP(ctx *gin.Context) {
    var inscription entities.Inscription
    if err := ctx.ShouldBindJSON(&inscription); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    // validacion de la inscripcion
    status, err := c.validateUseCase.Run(&inscription)
    if err != nil {
        if status == "rechazada" {
            ctx.JSON(http.StatusOK, gin.H{"status": "rechazada", "message": err.Error()})
        } else {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating inscription"})
        }
        return
    }

    // actualizacion en la bd
    if err := c.updateStatusUseCase.Run(&inscription, status); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating inscription status"})
        return
    }

    // respuesta
    ctx.JSON(http.StatusOK, gin.H{"status": status})
}

