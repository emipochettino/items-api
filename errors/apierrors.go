package errors

import (
	"strings"
	"net/http"
)

const (
	BadRequestMessage          = "Invalid request parameters."
	ResourceNotFoundMessage    = "Resource not found."
	InternalServerErrorMessage = "Internal server error."
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Initializes an APIError.
func newAPIError(code int, message string) *APIError {
	return &APIError{
		Status:  code,
		Message: message,
	}
}

func NewBadRequest(messages ...string) *APIError {
	message := BadRequestMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return newAPIError(http.StatusBadRequest, message)
}

func NewResourceNotFound(messages ...string) *APIError {
	message := ResourceNotFoundMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return newAPIError(http.StatusNotFound, message)
}

func NewInternalServerError(messages ...string) *APIError {
	message := InternalServerErrorMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return newAPIError(http.StatusInternalServerError, message)
}
