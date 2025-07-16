package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config interface {
	GetSlogConfig() Slog
	GetHttpServerConfig() HttpServer
}

type config struct {
	SlogConfig       *slog       `yaml:"slog"`
	HttpServerConfig *httpServer `yaml:"http_server"`
}

func NewConfig(configPath string) (Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}(file)

	config := &config{}

	return config, yaml.NewDecoder(file).Decode(config)
}

func (config *config) GetSlogConfig() Slog {
	return config.SlogConfig
}

func (config *config) GetHttpServerConfig() HttpServer {
	return config.HttpServerConfig
}
