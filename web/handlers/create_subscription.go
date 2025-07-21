package handlers

import (
	"encoding/json"
	"github.com/compico/em-task/internal/pkg/entity"
	"github.com/compico/em-task/internal/pkg/service"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/web/middleware"
	"github.com/compico/em-task/web/request"
	"net/http"
)

type CreateSubscriptionHandler struct {
	logger  logger.Logger
	service service.Subscription
}

func NewCreateSubscriptionHandler(logger logger.Logger, service service.Subscription) *CreateSubscriptionHandler {
	return &CreateSubscriptionHandler{
		logger:  logger,
		service: service,
	}
}

// CreateSubscriptionHandler создает новую подписку
// @Summary Создать подписку
// @Description Создает новую подписку в системе
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body request.CreateSubscriptionData true "Данные подписки"
// @Success 201 {object} SubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [post]
func (h *CreateSubscriptionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log, ok := ctx.Value(middleware.LoggerKey{}).(logger.Logger)
	if !ok {
		jsonError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	log.DebugContext(ctx, "New CreateSubscription request")

	req, err := request.CreateSubscriptionDataFromRequest(r)
	if err != nil {
		log.ErrorContext(ctx, "error on validate request", "error", err)
		jsonError(w, http.StatusBadRequest, "invalid request payload")

		return
	}
	log.DebugContext(ctx, "Parsed request", "request_data", req)

	sub := &entity.Subscription{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   req.GetStartDateAsTime(),
	}

	if err := h.service.Create(ctx, sub); err != nil {
		log.ErrorContext(
			ctx, "error on create subscription",
			"error", err,
			"payload", sub,
		)
		jsonError(w, http.StatusInternalServerError, "error on create subscription")

		return
	}

	log.Info("new subscription created", "subscription", sub)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&SubscriptionResponse{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate.Format("01-2006"),
	}); err != nil {
		log.ErrorContext(ctx, "error on encode response", "error", err)
		jsonError(w, http.StatusInternalServerError, "error on encode response")
	}
	log.DebugContext(ctx, "Handler successfully create subscription")
}
