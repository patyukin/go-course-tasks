package main

import (
	"context"
	"fmt"
	"github.com/patyukin/go-course-tasks/assesment_1/internal/config"
	"github.com/patyukin/go-course-tasks/assesment_1/internal/loader"
	"github.com/patyukin/go-course-tasks/assesment_1/internal/usecase"
	"github.com/patyukin/go-course-tasks/assesment_1/pkg/logger"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error creating config: %v", err)
	}

	l, err := logger.New(cfg)
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}

	errCh := make(chan error)
	ldr := loader.New(cfg, l)
	uc, err := usecase.New(ctx, ldr, cfg, l)
	if err != nil {
		slog.ErrorContext(ctx, "Error creating usecase: %v", err)
		os.Exit(1)
	}

	go uc.Consume(ctx, errCh)
	go uc.Produce(ctx, errCh)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errCh:
		l.ErrorContext(ctx, "Error consuming messages: %v", err)
	case res := <-sigChan:
		if res == syscall.SIGINT || res == syscall.SIGTERM {
			l.Info("Signal received", slog.String("signal", res.String()))
		} else if res == syscall.SIGHUP {
			l.Info("Signal received", slog.String("signal", res.String()))
		}
	}

	fmt.Println("StRShutting down...")
	cancel()

	fmt.Println("Server stopped.")
}
