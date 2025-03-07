package ports

type IValidationRepository interface {
    Validate(inscriptionID int) (string, error) 
    UpdateStatus(inscriptionID int, status string) error
}