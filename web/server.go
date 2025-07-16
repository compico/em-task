package web

import (
	"context"
	"net/http"
)

type Server interface {
	Start() error
	Stop(ctx context.Context) error
}

type server struct {
	httpServer *http.Server
}

func NewServer(httpServer *http.Server) Server {
	return &server{
		httpServer: httpServer,
	}
}

func (s *server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
