package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/compico/em-task/internal/pkg/service"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/web/middleware"
	"net/http"
	"strconv"
)

type DeleteSubscriptionResponse struct {
	ID int `json:"id" example:"123"`
} //@name DeleteSubscriptionResponse

type DeleteSubscriptionHandler struct {
	logger  logger.Logger
	service service.Subscription
}

func NewDeleteSubscriptionHandler(logger logger.Logger, service service.Subscription) *DeleteSubscriptionHandler {
	return &DeleteSubscriptionHandler{
		logger:  logger,
		service: service,
	}
}

// DeleteSubscriptionHandler удаляет подписку
// @Summary Удалить подписку
// @Description Удаляет подписку по её ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки" example:42
// @Success 204 {object} DeleteSubscriptionResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [delete]
func (h *DeleteSubscriptionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log, ok := ctx.Value(middleware.LoggerKey{}).(logger.Logger)
	if !ok {
		jsonError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	log.DebugContext(ctx, "New DeleteSubscription request")
	stringId := r.PathValue("id")
	if stringId == "" {
		log.ErrorContext(ctx, "empty subscription id", "error", fmt.Errorf("empty subscription id"))
		jsonError(w, http.StatusBadRequest, "invalid request")
		return
	}
	id, err := strconv.Atoi(stringId)
	if err != nil {
		log.ErrorContext(ctx, "error on parse subscription id", "error", err)
		jsonError(w, http.StatusBadRequest, "invalid subscription id")
		return
	}

	log.DebugContext(ctx, "Parsed request", "subscription_id", id)
	if err := h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, service.ErrResourceNotFound) {
			jsonError(w, http.StatusNotFound, "subscription not found")

			return
		}
		log.ErrorContext(ctx, "error on delete subscription", "error", err)
		jsonError(w, http.StatusInternalServerError, "error on delete subscription")

		return
	}

	w.WriteHeader(http.StatusNoContent)
	if err := json.NewEncoder(w).Encode(&DeleteSubscriptionResponse{
		ID: id,
	}); err != nil {
		log.ErrorContext(ctx, "error on encode response", "error", err)
		jsonError(w, http.StatusInternalServerError, "error on delete subscription")

		return
	}

	log.DebugContext(ctx, "Handler successfully delete subscription")
}
