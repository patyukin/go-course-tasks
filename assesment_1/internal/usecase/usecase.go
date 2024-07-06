package usecase

import (
	"context"
	"fmt"
	"github.com/patyukin/go-course-tasks/assesment_1/internal/config"
	"github.com/patyukin/go-course-tasks/assesment_1/internal/model"
	"log/slog"
	"sync"
)

type Loader interface {
	LoadMessages(ctx context.Context) ([]model.Message, error)
	LoadWhiteListTokens(ctx context.Context) (map[string]bool, error)
}

type UseCase struct {
	inputChs    []chan model.Message
	shutdownCh  chan struct{}
	shutdownWg  sync.WaitGroup
	messages    []model.Message
	cache       model.Cache
	validTokens map[string]bool
	mu          sync.Mutex
	cfg         *config.Config
	l           *slog.Logger
}

func New(ctx context.Context, ldr Loader, cfg *config.Config, l *slog.Logger) (*UseCase, error) {
	validTokens, err := ldr.LoadWhiteListTokens(ctx)

	if err != nil {
		return nil, fmt.Errorf("error loading tokens: %v", err)
	}

	messages, err := ldr.LoadMessages(ctx)
	if err != nil {
		return nil, fmt.Errorf("error loading messages: %v", err)
	}

	uc := &UseCase{
		mu:          sync.Mutex{},
		validTokens: validTokens,
		messages:    messages,
		cache:       make(model.Cache),
		inputChs:    make([]chan model.Message, cfg.InputChsCount),
		shutdownCh:  make(chan struct{}),
		cfg:         cfg,
		l:           l,
	}

	for i := range uc.inputChs {
		uc.inputChs[i] = make(chan model.Message, 100)
	}

	return uc, nil
}
