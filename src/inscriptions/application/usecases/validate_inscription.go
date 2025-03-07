package usecases

import (
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/domain/entities"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/domain/ports"
)

type ValidateInscriptionUseCase struct {
    validationRepo ports.IValidationRepository
}

func NewValidateInscriptionUseCase(validationRepo ports.IValidationRepository) *ValidateInscriptionUseCase {
    return &ValidateInscriptionUseCase{
        validationRepo: validationRepo,
    }
}

func (uc *ValidateInscriptionUseCase) Run(inscription *entities.Inscription) error {
    // valida inscripcion
    status, err := uc.validationRepo.Validate(inscription.ID)
    if err != nil {
        return err
    }

    //actualiza el stattus pa la bd
    inscription.Status = status

    
    return uc.validationRepo.UpdateStatus(inscription.ID, status)
}