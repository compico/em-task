package router

import (
	_ "github.com/compico/em-task/docs"
	"github.com/compico/em-task/web/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title Subscription API
// @version 1.0
// @description API для управления подписками
// @termsOfService http://swagger.io/terms/
// @host https://em-task.compico.su
// @BasePath /api/v1
// @schemes https

func NewServerMux(
	healthCheckHandler *handlers.HealthCheckHandler,
	subscriptionHandlers handlers.SubscriptionHandlers,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/subscriptions", subscriptionHandlers.SubscriptionsCreateHandler)
	mux.HandleFunc("GET /api/v1/subscriptions", subscriptionHandlers.SubscriptionsListHandler)
	mux.HandleFunc("GET /api/v1/subscriptions/{id}", subscriptionHandlers.SubscriptionsReadHandler)
	mux.HandleFunc("PUT /api/v1/subscriptions/{id}", subscriptionHandlers.SubscriptionsUpdateHandler)
	mux.HandleFunc("DELETE /api/v1/subscriptions/{id}", subscriptionHandlers.SubscriptionsDeleteHandler)

	// Swagger endpoint
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	mux.Handle("GET /health", healthCheckHandler)

	return mux
}
