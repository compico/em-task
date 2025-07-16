package di

import "github.com/compico/em-task/internal/pkg/config"

func SlogConfigProvider(config config.Config) config.Slog {
	return config.GetSlogConfig()
}

func HttpServerConfigProvider(config config.Config) config.HttpServer {
	return config.GetHttpServerConfig()
}
