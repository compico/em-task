package config

import slog2 "log/slog"

type Slog interface {
	GetAddSource() bool
	GetLevel() slog2.Level
}

type slog struct {
	AddSource bool   `yaml:"add_source"`
	Level     string `yaml:"log_level"`
}

func (config *slog) GetAddSource() bool {
	return config.AddSource
}

func (config *slog) GetLevel() slog2.Level {
	switch config.Level {
	case "debug":
		return slog2.LevelDebug
	case "info":
		return slog2.LevelInfo
	case "warn":
		return slog2.LevelWarn
	case "error":
		return slog2.LevelError
	default:
		return slog2.LevelInfo
	}
}
