package logger

import (
	"context"
	"fmt"
	"github.com/patyukin/go-course-tasks/assesment_1/internal/config"
	"log/slog"
	"os"
)

type MultiHandler struct {
	handlers []slog.Handler
}

func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}

	return false
}

func (m *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, h := range m.handlers {
		if err := h.Handle(ctx, record); err != nil {
			return err
		}
	}

	return nil
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var handlers []slog.Handler
	for _, h := range m.handlers {
		handlers = append(handlers, h.WithAttrs(attrs))
	}

	return NewMultiHandler(handlers...)
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
	var handlers []slog.Handler
	for _, h := range m.handlers {
		handlers = append(handlers, h.WithGroup(name))
	}

	return NewMultiHandler(handlers...)
}

func New(cfg *config.Config) (*slog.Logger, error) {
	var handlers []slog.Handler
	level := slog.LevelDebug
	if cfg.Env == "prod" {
		level = slog.LevelInfo
	}

	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	handlers = append(handlers, consoleHandler)

	if cfg.Env != "prod" {
		file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("error opening file: %v", err)
		}

		fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{Level: level})
		handlers = append(handlers, fileHandler)
	}

	multiHandler := NewMultiHandler(handlers...)

	logger := slog.New(multiHandler)

	return logger, nil
}
