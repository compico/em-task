package logger

import (
	"context"
	"log"
	"log/slog"
)

// Logger slog wrapper
type Logger interface {
	Debug(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	Info(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	Warn(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	Error(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
	With(args ...any) Logger
	WithGroup(name string) Logger
	GetStdLogger() *log.Logger
}

type logger struct {
	slog     *slog.Logger
	logLevel slog.Level
}

type OptionFunc func(*logger)

func NewLogger(logLevel slog.Level, s *slog.Logger, opts ...OptionFunc) Logger {
	l := &logger{
		logLevel: logLevel,
		slog:     s,
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func SetSlogDefault() OptionFunc {
	return func(l *logger) {
		l.SetSlogDefault()
	}
}

func (l *logger) With(args ...any) Logger {
	return &logger{
		logLevel: l.logLevel,
		slog:     l.slog.With(args...),
	}
}

func (l *logger) WithGroup(name string) Logger {
	return &logger{
		logLevel: l.logLevel,
		slog:     l.slog.WithGroup(name),
	}
}

func (l *logger) SetSlogDefault() {
	slog.SetDefault(l.slog)
}

func (l *logger) Debug(msg string, args ...any) {
	l.slog.Debug(msg, args...)
}

func (l *logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.slog.DebugContext(ctx, msg, args...)
}

func (l *logger) Info(msg string, args ...any) {
	l.slog.Info(msg, args...)
}

func (l *logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.slog.InfoContext(ctx, msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.slog.Warn(msg, args...)
}

func (l *logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.slog.WarnContext(ctx, msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.slog.Error(msg, args...)
}

func (l *logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.slog.ErrorContext(ctx, msg, args...)
}

func (l *logger) GetStdLogger() *log.Logger {
	return slog.NewLogLogger(l.slog.Handler(), l.logLevel)
}
