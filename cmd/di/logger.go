package di

import (
	"github.com/compico/em-task/internal/pkg/config"
	"github.com/compico/em-task/pkg/logger"
	"io"
	"log/slog"
	"os"
)

type SlogReplacerAttribute interface {
	Replace(groups []string, a slog.Attr) slog.Attr
}

func SlogProvider(handler slog.Handler) *slog.Logger {
	return slog.New(handler)
}

func SlogJsonHandlerProvider(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return slog.NewJSONHandler(w, opts)
}

func SlogJsonHandlerOptionsProvider(config config.Slog, slogReplacerAttr SlogReplacerAttribute) *slog.HandlerOptions {
	handleOpts := &slog.HandlerOptions{
		AddSource: config.GetAddSource(),
		Level:     config.GetLevel(),
	}

	if slogReplacerAttr != nil {
		handleOpts.ReplaceAttr = slogReplacerAttr.Replace
	}

	return handleOpts
}

func SlogWriterProvider() io.Writer {
	return os.Stdout
}

func SlogReplacerAttrProvider() SlogReplacerAttribute {
	return nil
}

func SlogLevelProvider(conf config.Slog) slog.Level {
	return conf.GetLevel()
}

func LoggerOptionsProvider() []logger.OptionFunc {
	return []logger.OptionFunc{
		logger.SetSlogDefault(),
	}
}
