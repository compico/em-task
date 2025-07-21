package router

import (
	_ "github.com/compico/em-task/docs"
	"github.com/compico/em-task/web/handlers"
	"github.com/compico/em-task/web/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Router struct {
	serve http.Handler
}

// @title Subscription API
// @version 1.0
// @description API для управления подписками
// @termsOfService http://swagger.io/terms/
// @host em-task.compico.su
// @BasePath /api/v1
// @schemes https

func NewServerMux(
	jsonResponseMiddleware *middleware.JsonResponse,
	withLoggerMiddleware *middleware.WithLogger,

	healthCheckHandler *handlers.HealthCheckHandler,
	createHandler *handlers.CreateSubscriptionHandler,
	readHandler *handlers.ReadSubscriptionHandler,
	updateHandler *handlers.UpdateSubscriptionHandler,
	deleteHandler *handlers.DeleteSubscriptionHandler,
	listHandler *handlers.ListSubscriptionsHandler,
	sumHandler *handlers.SumSubscriptionHandler,
) *Router {
	mux := http.NewServeMux()

	mux.Handle("POST /api/v1/subscriptions", jsonResponseMiddleware.Use(withLoggerMiddleware.Use(createHandler)))
	mux.Handle("GET /api/v1/subscriptions", jsonResponseMiddleware.Use(withLoggerMiddleware.Use(listHandler)))
	mux.Handle("GET /api/v1/subscriptions/{id}", jsonResponseMiddleware.Use(withLoggerMiddleware.Use(readHandler)))
	mux.Handle("PUT /api/v1/subscriptions/{id}", jsonResponseMiddleware.Use(withLoggerMiddleware.Use(updateHandler)))
	mux.Handle("DELETE /api/v1/subscriptions/{id}", jsonResponseMiddleware.Use(withLoggerMiddleware.Use(deleteHandler)))
	mux.Handle("GET /api/v1/subscriptions/sum", jsonResponseMiddleware.Use(withLoggerMiddleware.Use(sumHandler)))

	mux.Handle("GET /health", jsonResponseMiddleware.Use(healthCheckHandler))

	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	return &Router{
		serve: mux,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.serve.ServeHTTP(w, req)
}
