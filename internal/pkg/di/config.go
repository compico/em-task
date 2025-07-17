package di

import (
	"github.com/compico/em-task/internal/pkg/config"
	"github.com/compico/em-task/pkg/postgres"
)

func SlogConfigProvider(config config.Config) config.Slog {
	return config.GetSlogConfig()
}

func HttpServerConfigProvider(config config.Config) config.HttpServer {
	return config.GetHttpServerConfig()
}

func DatabaseConfigProvider(config config.Config) config.Database {
	return config.GetDatabaseConfig()
}

func ConnectionConfigProvider(dbConfig config.Database) postgres.ConnectionConfig {
	return dbConfig
}
