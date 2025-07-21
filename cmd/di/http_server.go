package di

import (
	"github.com/compico/em-task/internal/pkg/config"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/web/router"
	"net/http"
)

func HttpServerProvider(conf config.HttpServer, logger logger.Logger, handler *router.Router) *http.Server {
	return &http.Server{
		Addr:                         conf.GetAddr(),
		ReadTimeout:                  conf.GetReadTimeout(),
		ReadHeaderTimeout:            conf.GetReadHeaderTimeout(),
		WriteTimeout:                 conf.GetWriteTimeout(),
		IdleTimeout:                  conf.GetIdleTimeout(),
		MaxHeaderBytes:               conf.GetMaxHeaderBytes(),
		DisableGeneralOptionsHandler: conf.GetDisableGeneralOptionsHandler(),
		Handler:                      handler,
		ErrorLog:                     logger.GetStdLogger(),
	}
}
