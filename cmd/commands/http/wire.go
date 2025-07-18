//go:build wireinject
// +build wireinject

package http

import (
	"context"
	"github.com/compico/em-task/cmd/di"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/pkg/postgres"
	"github.com/compico/em-task/web"
	"github.com/google/wire"
)

type (
	App struct {
		server web.Server
		logger logger.Logger
		pg     postgres.DB
	}
)

func InitializeApp(
	ctx context.Context,
	filepath string,
) (*App, func(), error) {
	panic(wire.Build(
		di.BaseSet,
		di.LoggerSet,
		di.DatabaseSet,
		di.HttpServerSet,

		wire.Struct(new(App), "*"),
	))
}
