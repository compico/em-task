package handlers

import (
	"encoding/json"
	"net/http"
)

func jsonError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request"`
} //@name ErrorResponse
