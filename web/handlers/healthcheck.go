package handlers

import (
	"encoding/json"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/pkg/postgres"
	"net/http"
	"time"
)

type HealthCheckHandler struct {
	logger    logger.Logger
	startTime time.Time
	db        postgres.DB
}

func NewHealthCheck(logger logger.Logger, db postgres.DB) *HealthCheckHandler {
	return &HealthCheckHandler{
		logger:    logger,
		startTime: time.Now(),
		db:        db,
	}
}

type healthCheckResponseBody struct {
	Status   string `json:"status"`
	Database struct {
		Status string `json:"status"`
		Error  string `json:"error,omitempty"`
	} `json:"database"`
	Uptime string `json:"uptime"`
}

func (h *HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := healthCheckResponseBody{
		Status: "ok",
		Uptime: time.Since(h.startTime).Truncate(1 * time.Second).String(),
	}

	if err := h.db.Ping(r.Context()); err != nil {
		resp.Database.Error = "Error connecting to database"
		resp.Database.Status = "error"
	} else {
		resp.Database.Status = "ok"
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.ErrorContext(r.Context(), "error on encoding response:", "error", err)
	}
}
