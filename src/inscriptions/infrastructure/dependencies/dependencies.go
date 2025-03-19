package dependencies

import (
	"github.com/Andresito126/go-validation_inscriptions/src/core/database"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/usecases"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/adapters"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/controllers"
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/services"
)

func SetupValidationDependencies() (*controllers.ValidateInscriptionController, error) {
	
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	
	mysqlRepo := adapters.NewMySQLAdapter(db)

	// crear los casos de uso
	validateUseCase := usecases.NewValidateInscriptionUseCase(mysqlRepo)
	updateStatusUseCase := usecases.NewUpdateInscriptionStatusUseCase(mysqlRepo)

	// crear el rabbit adapter
	rabbit := adapters.NewRabbit()

	// eventos y que pasa el adaptador Rabbit
	eventService := services.NewEventService(rabbit)

	// controlador de validación de inscripción
	return controllers.NewValidateInscriptionController(validateUseCase, updateStatusUseCase, eventService), nil
}
