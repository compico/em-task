//go:build wireinject
// +build wireinject

package di

import (
	"github.com/compico/em-task/internal/pkg/config"
	"github.com/compico/em-task/internal/pkg/pgrepo"
	"github.com/compico/em-task/internal/pkg/service"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/pkg/postgres"
	"github.com/compico/em-task/web"
	"github.com/compico/em-task/web/handlers"
	"github.com/compico/em-task/web/middleware"
	"github.com/compico/em-task/web/router"
	"github.com/google/wire"
)

var BaseSet = wire.NewSet(
	config.NewConfig,
)

var SubscriptionHandlersSet = wire.NewSet(
	handlers.NewCreateSubscriptionHandler,
	handlers.NewReadSubscriptionHandler,
	handlers.NewUpdateSubscriptionHandler,
	handlers.NewDeleteSubscriptionHandler,
	handlers.NewListSubscriptionsHandler,
	handlers.NewSumSubscriptionHandler,
)

var MiddlewaresSet = wire.NewSet(
	middleware.NewJsonResponseMiddleware,
	middleware.NewWithLoggerMiddleware,
)

var HttpHandlersSet = wire.NewSet(
	MiddlewaresSet,
	handlers.NewHealthCheck,
	SubscriptionHandlersSet,
)

var SubscriptionSet = wire.NewSet(
	pgrepo.NewSubscriptionRepository,
	service.NewSubscription,
)

var HttpServerSet = wire.NewSet(
	HttpServerConfigProvider,
	HttpHandlersSet,
	router.NewServerMux,
	HttpServerProvider,
	web.NewServer,
)

var LoggerSet = wire.NewSet(
	SlogConfigProvider,
	SlogWriterProvider,
	SlogJsonHandlerOptionsProvider,
	SlogReplacerAttrProvider,
	SlogJsonHandlerProvider,
	SlogProvider,
	SlogLevelProvider,
	LoggerOptionsProvider,

	logger.NewLogger,
)

var DatabaseSet = wire.NewSet(
	DatabaseConfigProvider,
	ConnectionConfigProvider,
	postgres.NewConnection,
	postgres.NewDatabase,
)
