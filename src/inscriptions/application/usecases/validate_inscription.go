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

func (uc *ValidateInscriptionUseCase) Run(inscription *entities.Inscription) (string, error) {
    status, err := uc.validationRepo.Validate(inscription.ID)
    if err != nil {
        return "", err
    }

    return status, nil
}
