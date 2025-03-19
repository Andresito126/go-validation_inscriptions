package services

import "github.com/Andresito126/go-validation_inscriptions/src/inscriptions/application/repositories"

type EventService struct {
    rabbit repositories.IRabbitRepository
}

func NewEventService(rabbit repositories.IRabbitRepository) *EventService {
    return &EventService{rabbit: rabbit}
}

func (e *EventService) PublishNotification(studentID, courseID, status string) {
    e.rabbit.SendMessageToBroker(studentID, courseID, status)
}
