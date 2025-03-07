package main

import (
    "log"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/dependencies"
)

func main() {
   
    controller, err := dependencies.SetupValidationDependencies()
    if err != nil {
        log.Fatalf("Error al configurar dependencias: %v", err)
    }

    // consumo de msg
    log.Println("API 2 iniciada")
    controller.Start() 
}