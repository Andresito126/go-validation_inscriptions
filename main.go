package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/dependencies"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/routes"
)

func main() {

    router := gin.Default()

    validateInscriptionController, err := dependencies.SetupValidationDependencies()
    if err != nil {
        log.Fatalf("Error al configurar las dependencias: %v", err)
    }

    routes.InscriptionRoutes(router, validateInscriptionController)


    if err := router.Run(":8081"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}
