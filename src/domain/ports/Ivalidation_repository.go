package ports

type IValidationRepository interface {
    Validate(inscriptionID int) (string, error) 
}