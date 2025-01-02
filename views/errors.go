package views

import (
	"encoding/json"
	"net/http"
)

// APIError representa um erro padronizado na API.
type APIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

// NewAPIError cria uma nova instância de APIError.
func NewAPIError(message string, statusCode int) *APIError {
	return &APIError{
		Message:    message,
		StatusCode: statusCode,
	}
}

// RenderError envia um erro padronizado como resposta JSON.
func RenderError(w http.ResponseWriter, err *APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(err)
}
