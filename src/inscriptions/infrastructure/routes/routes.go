package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/controllers"
)

func InscriptionRoutes(router *gin.Engine, validateInscriptionController *controllers.ValidateInscriptionController) {
    routes := router.Group("/inscriptions")

    routes.POST("/validate", validateInscriptionController.HandleHTTP)
}
