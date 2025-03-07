package dependencies

import (
    "github.com/Andresito126/go-validation_inscriptions/src/core/database"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/services"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/usecases"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/adapters"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/infrastructure/controllers"
)

func SetupValidationDependencies() (*controllers.ValidateInscriptionController, error) {
    // conn mysql
    db, err := database.Connect()
    if err != nil {
        return nil, err
    }

    // mysq adp
    mysqlRepo := adapters.NewMySQLAdapter(db)

    //conn rabbit
    rabbitRepo, err := adapters.NewRabbitAdapter()
    if err != nil {
        return nil, err
    }

    // inicia rabbit
    rabbitService := services.NewRabbitService(rabbitRepo)

    
    useCase := usecases.NewValidateInscriptionUseCase(mysqlRepo)

    return controllers.NewValidateInscriptionController(useCase, rabbitService), nil
}