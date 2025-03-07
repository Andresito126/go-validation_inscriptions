// src/api2/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"github.com/Andresito126/go-validation_inscriptions/src/core/database"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/services"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/usecases"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/adapters"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/controllers"
)

func SetupValidationDependencies() (*controllers.ValidateInscriptionController, error) {
    
    db, err := database.Connect()
    if err != nil {
        return nil, err
    }

    
    mysqlRepo := adapters.NewMySQLAdapter(db)

    
    rabbitRepo, err := adapters.NewRabbitAdapter()
    if err != nil {
        return nil, err
    }

 //servicio rab
 rabbitService := services.NewRabbitService(rabbitRepo)

 //uc
 useCase := usecases.NewValidateInscriptionUseCase(mysqlRepo)

 //controller
 return controllers.NewValidateInscriptionController(useCase, rabbitService), nil
}