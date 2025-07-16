//go:build wireinject
// +build wireinject

package http

import (
	"context"
	"github.com/compico/em-task/internal/pkg/config"
	"github.com/compico/em-task/internal/pkg/di"
	"github.com/compico/em-task/internal/pkg/logger"
	"github.com/compico/em-task/web"
	"github.com/google/wire"
)

type (
	App struct {
		server web.Server
		logger logger.Logger
	}
)

func InitializeApp(
	ctx context.Context,
	filepath string,
) (*App, error) {
	panic(wire.Build(
		config.NewConfig,

		di.LoggerSet,
		di.HttpServerSet,

		wire.Struct(new(App), "*"),
	))
}
