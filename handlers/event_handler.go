package handlers

import (
	"encoding/json"
	"net/http"
	"server/services"
	"server/utils"
)

// EventHandler will respond to request related to events.
type EventHandler interface {
	// GetEvents will return all events.
	GetEvents() http.HandlerFunc
}

// DefaultEventHandler is default implementation of [EventHandler].
type DefaultEventHandler struct {
	eventRepository services.EventService
}

func (h *DefaultEventHandler) GetEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, errorResponse := h.eventRepository.GetEvents(r.Context())
		if errorResponse != nil {
			utils.RespondWithError(w, errorResponse)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(events)
		if err != nil {
			utils.RespondWithError(w, utils.InternalServerError())
		}
	}
}

// NewDefaultEventHandler will create new [DefaultEventHandler].
func NewDefaultEventHandler(eventRepository services.EventService) *DefaultEventHandler {
	return &DefaultEventHandler{
		eventRepository: eventRepository,
	}
}
