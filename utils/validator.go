package utils

import "net/http"

type ValidatablePayload interface {
	Validate() *ErrorResponse
}

func CheckPayload(w http.ResponseWriter, payload ValidatablePayload) bool {
	errorResponse := payload.Validate()
	if errorResponse != nil {
		RespondWithError(w, *errorResponse)
		return true
	}
	return false
}
