package handlers

import (
	"encoding/json"
	"github.com/compico/em-task/internal/pkg/filter"
	"github.com/compico/em-task/internal/pkg/service"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/web/middleware"
	"net/http"
)

type SumSubscriptionResponse struct {
	Sum int64 `json:"sum"`
} //@name SumSubscriptionResponse

type SumSubscriptionHandler struct {
	logger  logger.Logger
	service service.Subscription
}

func NewSumSubscriptionHandler(logger logger.Logger, service service.Subscription) *SumSubscriptionHandler {
	return &SumSubscriptionHandler{
		logger:  logger,
		service: service,
	}
}

// SumSubscriptionHandler Получение суммы цен с фильтрацией
// @Summary Получить сумму цен подписок
// @Description Получает сумму цены с возможностью фильтрации по различным параметрам
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param from query string false "Фильтр по начальной дате (формат: MM-YYYY)" example:"01-2024"
// @Param to query string false "Фильтр по конечной дате (формат: MM-YYYY)" example:"12-2025"
// @Param user_id query string false "Фильтр по ID пользователя (UUID)" example:"550e8400-e29b-41d4-a716-446655440000"
// @Param service_name query string false "Фильтр по названию сервиса" example:"Yandex Plus"
// @Success 200 {object} SumSubscriptionResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/sum [get]
func (h *SumSubscriptionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log, ok := ctx.Value(middleware.LoggerKey{}).(logger.Logger)
	if !ok {
		jsonError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	log.Debug("New sum subscription request")
	f, err := filter.ListSubscriptionsFromQuery(r.URL.Query())
	if err != nil {
		log.ErrorContext(ctx, "error on read subscriptions", "error", err)
		jsonError(w, http.StatusInternalServerError, "internal server error")

		return
	}

	log.DebugContext(ctx, "read filter params", "subscription_filter", f)
	sum, err := h.service.Sum(ctx, f)
	if err != nil {
		log.ErrorContext(ctx, "error on get sum", "error", err)
		jsonError(w, http.StatusInternalServerError, "internal server error")

		return
	}

	log.DebugContext(ctx, "got sum", "sum", sum)
	if err := json.NewEncoder(w).Encode(sum); err != nil {
		log.ErrorContext(ctx, "error on encode sum", "error", err)
		jsonError(w, http.StatusInternalServerError, "internal server error")

		return
	}

	log.InfoContext(ctx, "Handler successfully sum subscriptions")
}
