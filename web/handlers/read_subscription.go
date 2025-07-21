package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/compico/em-task/internal/pkg/service"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/web/middleware"
	"net/http"
	"strconv"
)

type ReadSubscriptionHandler struct {
	logger  logger.Logger
	service service.Subscription
}

func NewReadSubscriptionHandler(logger logger.Logger, service service.Subscription) *ReadSubscriptionHandler {
	return &ReadSubscriptionHandler{
		logger:  logger,
		service: service,
	}
}

// ReadSubscriptionHandler получает подписку по ID
// @Summary Получить подписку
// @Description Получает подписку по её ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки" example:42
// @Success 200 {object} SubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [get]
func (h *ReadSubscriptionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log, ok := ctx.Value(middleware.LoggerKey{}).(logger.Logger)
	if !ok {
		jsonError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	log.DebugContext(ctx, "New ReadSubscription request")
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

	log.DebugContext(ctx, "Parsed subscription id", "subscription_id", id)

	sub, err := h.service.Read(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			jsonError(w, http.StatusNotFound, "subscription not found")

			return
		}
		log.ErrorContext(ctx, "error on read subscription", "error", err)
		jsonError(w, http.StatusInternalServerError, "error on read subscription")

		return
	}

	if sub == nil {
		jsonError(w, http.StatusNotFound, "subscription not found")
		return
	}

	if err := json.NewEncoder(w).Encode(&SubscriptionResponse{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate.Format("01-2006"),
	}); err != nil {
		log.ErrorContext(ctx, "error on encode response", "error", err)
	}
	log.DebugContext(ctx, "Handler successfully read subscription")
}
