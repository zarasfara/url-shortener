package handlers

import (
	"encoding/json"
	"net/http"
)

// Custom error structure
type HttpError struct {
	Message string `json:"message"`
}

// NewHttpError create a new http error
func NewHttpError(message string) *HttpError {
	return &HttpError{
		Message: message,
	}
}

// Method to write the HttpError as a JSON response
func (e *HttpError) Write(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(e)
}
