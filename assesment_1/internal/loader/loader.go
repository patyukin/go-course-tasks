package loader

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/patyukin/go-course-tasks/assesment_1/internal/config"
	"github.com/patyukin/go-course-tasks/assesment_1/internal/model"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"sync"
)

type Data struct {
	Tokens []string `yaml:"tokens"`
	Files  []string `yaml:"files"`
}

type Loader struct {
	mu  sync.Mutex
	l   *slog.Logger
	cfg *config.Config
}

func New(cfg *config.Config, l *slog.Logger) *Loader {
	return &Loader{
		l:   l,
		cfg: cfg,
	}
}

func (l *Loader) LoadMessages(ctx context.Context) ([]model.Message, error) {
	l.l.InfoContext(ctx, "Loading messages")
	data, err := os.ReadFile("seed/data.yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var cfgData Data
	if err = yaml.Unmarshal(data, &cfgData); err != nil {
		return nil, fmt.Errorf("error unmarshalling file: %v", err)
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	msgs := make([]model.Message, 0, len(cfgData.Tokens)*len(cfgData.Files))
	for i, token := range cfgData.Tokens {
		for j, file := range cfgData.Files {
			msg := model.Message{
				Token:  token,
				FileID: file,
				Data:   fmt.Sprintf("Message %d%d from token %s and file %s", i, j, token, file),
			}

			l.l.InfoContext(ctx, "Loading messages: ", slog.Any("msg", msg))
			msgs = append(msgs, msg)
		}
	}

	return msgs, nil
}

func (l *Loader) LoadWhiteListTokens(ctx context.Context) (map[string]bool, error) {
	l.l.InfoContext(ctx, "Loading messages")
	file, err := os.Open("seed/valid_tokens.json")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	var validTokens []string
	if err = json.NewDecoder(file).Decode(&validTokens); err != nil {
		return nil, fmt.Errorf("error decoding file: %v", err)
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	tokens := make(map[string]bool, len(validTokens))
	for _, token := range validTokens {
		tokens[token] = true
	}

	l.l.InfoContext(ctx, "Loading messages: ", slog.Any("tokens", tokens))
	return tokens, nil
}
