//go:build wireinject
// +build wireinject

package di

import (
	"github.com/compico/em-task/internal/pkg/config"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/pkg/postgres"
	"github.com/compico/em-task/web"
	"github.com/compico/em-task/web/handlers"
	"github.com/compico/em-task/web/router"
	"github.com/google/wire"
)

var BaseSet = wire.NewSet(
	config.NewConfig,
	SlogConfigProvider,
)

var HttpHandlersSet = wire.NewSet(
	handlers.NewHealthCheck,
	handlers.NewSubscriptionHandler,
)

var HttpServerSet = wire.NewSet(
	HttpServerConfigProvider,
	HttpHandlersSet,
	router.NewServerMux,
	HttpServerProvider,
	web.NewServer,
)

var LoggerSet = wire.NewSet(
	SlogWriterProvider,
	SlogJsonHandlerOptionsProvider,
	SlogReplacerAttrProvider,
	SlogJsonHandlerProvider,
	SlogProvider,
	SlogLevelProvider,

	logger.NewLogger,
)

var DatabaseSet = wire.NewSet(
	DatabaseConfigProvider,
	ConnectionConfigProvider,
	postgres.NewConnection,
	postgres.NewDatabase,
)
