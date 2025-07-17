package handlers

import "net/http"

type CreateSubscriptionRequest struct {
	ServiceName string `json:"service_name" example:"Netflix" binding:"required"`
	Price       int    `json:"price" example:"999" binding:"required"`
	UserID      string `json:"user_id" example:"user123" binding:"required"`
	StartDate   string `json:"start_date" example:"01-2006" binding:"required"`
}

type UpdateSubscriptionRequest struct {
	ServiceName string `json:"service_name" example:"Netflix"`
	Price       int    `json:"price" example:"999"`
	UserID      string `json:"user_id" example:"user123"`
	StartDate   string `json:"start_date" example:"01-2006"`
}

type SubscriptionResponse struct {
	ID          int    `json:"id" example:"1"`
	ServiceName string `json:"service_name" example:"Netflix"`
	Price       int    `json:"price" example:"999"`
	UserID      string `json:"user_id" example:"user123"`
	StartDate   string `json:"start_date" example:"01-2006"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request"`
}

type SubscriptionHandlers interface {
	SubscriptionsCreateHandler(w http.ResponseWriter, r *http.Request)
	SubscriptionsReadHandler(w http.ResponseWriter, r *http.Request)
	SubscriptionsUpdateHandler(w http.ResponseWriter, r *http.Request)
	SubscriptionsDeleteHandler(w http.ResponseWriter, r *http.Request)
	SubscriptionsListHandler(w http.ResponseWriter, r *http.Request)
}

type subscriptionHandler struct {
}

func NewSubscriptionHandler() SubscriptionHandlers {
	return &subscriptionHandler{}
}

// SubscriptionsCreateHandler создает новую подписку
// @Summary Создать подписку
// @Description Создает новую подписку в системе
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body CreateSubscriptionRequest true "Данные подписки"
// @Success 201 {object} SubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [post]
func (s *subscriptionHandler) SubscriptionsCreateHandler(w http.ResponseWriter, r *http.Request) {
}

// SubscriptionsReadHandler получает подписку по ID
// @Summary Получить подписку
// @Description Получает подписку по её ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки"
// @Success 200 {object} SubscriptionResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [get]
func (s *subscriptionHandler) SubscriptionsReadHandler(w http.ResponseWriter, r *http.Request) {
}

// SubscriptionsUpdateHandler обновляет подписку
// @Summary Обновить подписку
// @Description Обновляет существующую подписку по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки"
// @Param subscription body UpdateSubscriptionRequest true "Обновленные данные подписки"
// @Success 200 {object} SubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [put]
func (s *subscriptionHandler) SubscriptionsUpdateHandler(w http.ResponseWriter, r *http.Request) {
}

// SubscriptionsDeleteHandler удаляет подписку
// @Summary Удалить подписку
// @Description Удаляет подписку по её ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки"
// @Success 204 "Подписка успешно удалена"
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [delete]
func (s *subscriptionHandler) SubscriptionsDeleteHandler(w http.ResponseWriter, r *http.Request) {
}

// SubscriptionsListHandler получает список подписок с фильтрацией
// @Summary Получить список подписок
// @Description Получает список подписок с возможностью фильтрации по различным параметрам
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param from query string false "Фильтр по начальной дате (формат: MM-YYYY)" example:"01-2006"
// @Param to query string false "Фильтр по конечной дате (формат: MM-YYYY)" example:"01-2006"
// @Param user_id query string false "Фильтр по ID пользователя" example:"user123"
// @Param service_name query string false "Фильтр по названию сервиса" example:"Netflix"
// @Success 200 {array} SubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [get]
func (s *subscriptionHandler) SubscriptionsListHandler(w http.ResponseWriter, r *http.Request) {
}
