package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/compico/em-task/internal/pkg/filter"
	"github.com/compico/em-task/internal/pkg/service"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/web/middleware"
	"github.com/compico/em-task/web/request"
	"net/http"
	"strconv"
)

type UpdateSubscriptionHandler struct {
	logger  logger.Logger
	service service.Subscription
}

func NewUpdateSubscriptionHandler(logger logger.Logger, service service.Subscription) *UpdateSubscriptionHandler {
	return &UpdateSubscriptionHandler{
		logger:  logger,
		service: service,
	}
}

// UpdateSubscriptionHandler обновляет подписку
// @Summary Обновить подписку
// @Description Обновляет существующую подписку по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки" example:42
// @Param subscription body request.UpdateSubscriptionData true "Обновленные данные подписки"
// @Success 200 {object} SubscriptionResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [put]
func (h *UpdateSubscriptionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log, ok := ctx.Value(middleware.LoggerKey{}).(logger.Logger)
	if !ok {
		jsonError(w, http.StatusInternalServerError, "internal server error")
		return
	}

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

	log.DebugContext(ctx, "New UpdateSubscription request", "subscription_id", id)
	req, err := request.UpdateSubscriptionDataFromRequest(r)
	if err != nil {
		log.ErrorContext(ctx, "error on validate request", "error", err)
		jsonError(w, http.StatusBadRequest, "invalid request payload")

		return
	}

	log.DebugContext(ctx, "Parsed request", "request_data", req)

	sub, err := h.service.Update(ctx, id, filter.UpdateSubscription{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   req.GetStartDate(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			jsonError(w, http.StatusNotFound, "subscription not found")
			return
		}
		log.ErrorContext(ctx, "error on update subscription", "error", err)
		jsonError(w, http.StatusInternalServerError, "error on update subscription")

		return
	}

	log.InfoContext(ctx, "Updated subscription", "subscription", sub)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&SubscriptionResponse{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate.Format("01-2006"),
	}); err != nil {
		log.ErrorContext(ctx, "error on encode response", "error", err)
		jsonError(w, http.StatusInternalServerError, "error on update subscription")

		return
	}

	log.DebugContext(ctx, "Handler successfully updated subscription")
}
