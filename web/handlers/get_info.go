package handlers

import (
	"encoding/json"
	"github.com/compico/em-task/internal/pkg/logger"
	"net/http"
)

type GetInfo struct {
	logger logger.Logger
}

func NewGetInfo(logger logger.Logger) *GetInfo {
	return &GetInfo{
		logger: logger,
	}
}

type responseBody struct {
	Status string `json:"status"`
}

func (h *GetInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := responseBody{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.ErrorContext(r.Context(), "error on encoding response:", "error", err)
	}
}
