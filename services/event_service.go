package services

import (
	"context"
	"net/http"
	"server/models"
	"server/repositories"
	"server/utils"
)

// EventService will manage events business logic.
type EventService interface {
	// GetEvents will return all events.
	GetEvents(ctx context.Context) ([]models.Event, *utils.ErrorResponse)

	// RegisterForEvent will register the user for an event.
	RegisterForEvent(ctx context.Context, userId, eventId int) *utils.ErrorResponse
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

func (s *DefaultEventService) RegisterForEvent(ctx context.Context, userId, eventId int) *utils.ErrorResponse {
	result, err := s.eventRepository.CheckEventId(ctx, eventId)
	if err != nil {
		return utils.InternalServerError()
	}

	if !result {
		return utils.NewErrorResponse("Invalid id", http.StatusBadRequest)
	}

	result, err = s.eventRepository.CheckIfRegistrationExists(ctx, userId, eventId)
	if err != nil {
		return utils.InternalServerError()
	}
	if result {
		return utils.NewErrorResponse("Already registered", http.StatusBadRequest)
	}

	err = s.eventRepository.AddRegistration(ctx, userId, eventId)
	if err != nil {
		return utils.InternalServerError()
	}

	return nil
}

// NewDefaultEventService will create new [DefaultEventService].
func NewDefaultEventService(eventRepository repositories.EventRepository) *DefaultEventService {
	return &DefaultEventService{
		eventRepository: eventRepository,
	}
}
