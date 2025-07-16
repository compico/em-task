package router

import (
	"github.com/compico/em-task/web/handlers"
	"net/http"
)

func NewServerMux(
	getInfo *handlers.GetInfo,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", http.NotFoundHandler().ServeHTTP)

	mux.Handle("GET /info", getInfo)

	return mux
}
