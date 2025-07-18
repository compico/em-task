package di

import (
	"github.com/compico/em-task/internal/pkg/config"
	"github.com/compico/em-task/pkg/logger"
	"net/http"
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
