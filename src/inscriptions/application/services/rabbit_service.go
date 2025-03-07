package services

import (
	"github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/repository"
	
)

type RabbitService struct {
    rabbitRepo repository.IRabbitService 
}

func NewRabbitService(rabbitRepo repository.IRabbitService) *RabbitService {
    return &RabbitService{
        rabbitRepo: rabbitRepo,
    }
}

func (s *RabbitService) Consume(queue string, handler func([]byte)) error {
    return s.rabbitRepo.Consume(queue, handler)
}