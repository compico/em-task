package handlers

import (
	"encoding/json"
	"github.com/compico/em-task/internal/pkg/service"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/web/middleware"
	"net/http"
	"strconv"
)

type ListSubscriptionsHandler struct {
	logger  logger.Logger
	service service.Subscription
}

func NewListSubscriptionsHandler(logger logger.Logger, service service.Subscription) *ListSubscriptionsHandler {
	return &ListSubscriptionsHandler{
		logger:  logger,
		service: service,
	}
}

// ListSubscriptionsHandler получает список подписок
// @Summary Получить список подписок
// @Description Получает список подписок
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param page query string false "Страница" example:"1"
// @Param per_page query string false "Элементов на страницу" example:"25"
// @Success 200 {array} SubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [get]
func (h *ListSubscriptionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log, ok := ctx.Value(middleware.LoggerKey{}).(logger.Logger)
	if !ok {
		jsonError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	log.DebugContext(ctx, "New ListSubscriptions request")

	page := 1
	if r.URL.Query().Has("page") {
		var err error
		page, err = strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			log.ErrorContext(ctx, "error on parse page", "error", err)
			jsonError(w, http.StatusBadRequest, "invalid request")

			return
		}
	}

	perPage := 25
	if r.URL.Query().Has("per_page") {
		var err error
		perPage, err = strconv.Atoi(r.URL.Query().Get("per_page"))
		if err != nil {
			log.ErrorContext(ctx, "error on parse per_page", "error", err)
			jsonError(w, http.StatusBadRequest, "invalid request")

			return
		}
	}

	l, err := h.service.List(ctx, page, perPage)
	if err != nil {
		log.ErrorContext(ctx, "error on list subscriptions", "error", err)
		jsonError(w, http.StatusInternalServerError, "error on list subscriptions")

		return
	}
	if len(l) == 0 {
		jsonError(w, http.StatusNotFound, "subscriptions not found")
		return
	}

	log.DebugContext(ctx, "Got subscriptions", "subscriptions_count", len(l))

	w.WriteHeader(http.StatusOK)
	var resp []*SubscriptionResponse
	for _, e := range l {
		resp = append(resp, &SubscriptionResponse{
			ID:          e.ID,
			ServiceName: e.ServiceName,
			Price:       e.Price,
			UserID:      e.UserID,
			StartDate:   e.StartDate.Format("01-2006"),
		})
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.ErrorContext(ctx, "error on encode response", "error", err)
		jsonError(w, http.StatusInternalServerError, "error on list subscriptions")

		return
	}

	log.DebugContext(ctx, "Handler successfully list subscriptions", "subscriptions_count", len(resp))
}
