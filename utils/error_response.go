package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

func NewErrorResponse(message string, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}

func RespondWithError(w http.ResponseWriter, response ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "There was an unknown error", http.StatusInternalServerError)
	}
}
