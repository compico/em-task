package middleware

import (
	"context"
	"github.com/compico/em-task/pkg/logger"
	"net/http"
)

type LoggerKey struct{}

type WithLogger struct {
	logger logger.Logger
}

func NewWithLoggerMiddleware(l logger.Logger) *WithLogger {
	return &WithLogger{
		logger: l,
	}
}

func (m *WithLogger) Use(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpLogger := m.logger.With(
			"endpoint", r.RequestURI,
			"method", r.Method,
			"remote_addr", r.Header.Get("X-Real-IP"),
			"user_agent", r.UserAgent(),
			"cf_ray", r.Header.Get("CF-Ray"),
		)

		ctx := context.WithValue(r.Context(), LoggerKey{}, httpLogger)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
