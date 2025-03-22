package services

import (
	"context"
	"server/models"
	"server/repositories"
	"server/utils"
)

// EventService will manage events business logic.
type EventService interface {
	GetEvents(ctx context.Context) ([]models.Event, *utils.ErrorResponse)
}

// DefaultEventService is the default implementation of [EventService].
type DefaultEventService struct {
	eventRepository repositories.EventRepository
}

func (s *DefaultEventService) GetEvents(ctx context.Context) ([]models.Event, *utils.ErrorResponse) {
	events, err := s.eventRepository.GetEvents(ctx)
	if err != nil {
		return nil, utils.InternalServerError()
	}
	return events, nil
}

// NewDefaultEventService will create new [DefaultEventService].
func NewDefaultEventService(eventRepository repositories.EventRepository) *DefaultEventService {
	return &DefaultEventService{
		eventRepository: eventRepository,
	}
}
