    package dependencies

    import (
        "github.com/Andresito126/go-validation_inscriptions/src/core/database"
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

        validateUseCase := usecases.NewValidateInscriptionUseCase(mysqlRepo)
        updateStatusUseCase := usecases.NewUpdateInscriptionStatusUseCase(mysqlRepo)

        return controllers.NewValidateInscriptionController(validateUseCase, updateStatusUseCase), nil
    }
