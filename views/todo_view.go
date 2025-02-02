package views

import (
	"encoding/json"
	"net/http"
)
// Usa RenderJSON para enviar a resposta
func RenderJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
