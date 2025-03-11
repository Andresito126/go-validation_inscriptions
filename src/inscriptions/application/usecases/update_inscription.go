package usecases

import (
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/domain/entities"
    "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/domain/ports"
)

type UpdateInscriptionStatusUseCase struct {
    validationRepo ports.IValidationRepository
}

func NewUpdateInscriptionStatusUseCase(validationRepo ports.IValidationRepository) *UpdateInscriptionStatusUseCase {
    return &UpdateInscriptionStatusUseCase{
        validationRepo: validationRepo,
    }
}

func (uc *UpdateInscriptionStatusUseCase) Run(inscription *entities.Inscription, status string) error {
    inscription.Status = status
    return uc.validationRepo.UpdateStatus(inscription.ID, status)
}
