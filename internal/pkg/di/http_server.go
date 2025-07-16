package di

import (
	"github.com/compico/em-task/internal/pkg/config"
	"github.com/compico/em-task/internal/pkg/logger"
	"github.com/compico/em-task/web"
	"github.com/compico/em-task/web/handlers"
	"github.com/compico/em-task/web/router"
	"github.com/google/wire"
	"net/http"
)

var HttpHandlersSet = wire.NewSet(
	handlers.NewGetInfo,
)

var HttpServerSet = wire.NewSet(
	HttpServerConfigProvider,
	HttpHandlersSet,
	router.NewServerMux,
	HttpServerProvider,
	web.NewServer,
)

func HttpServerProvider(conf config.HttpServer, logger logger.Logger, mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:                         conf.GetAddr(),
		ReadTimeout:                  conf.GetReadTimeout(),
		ReadHeaderTimeout:            conf.GetReadHeaderTimeout(),
		WriteTimeout:                 conf.GetWriteTimeout(),
		IdleTimeout:                  conf.GetIdleTimeout(),
		MaxHeaderBytes:               conf.GetMaxHeaderBytes(),
		DisableGeneralOptionsHandler: conf.GetDisableGeneralOptionsHandler(),
		Handler:                      mux,
		ErrorLog:                     logger.GetStdLogger(),
	}
}
