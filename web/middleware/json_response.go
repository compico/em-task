package middleware

import "net/http"

type JsonResponse struct{}

func NewJsonResponseMiddleware() *JsonResponse {
	return &JsonResponse{}
}

func (m *JsonResponse) Use(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
