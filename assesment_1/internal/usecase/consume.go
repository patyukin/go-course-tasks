package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/patyukin/go-course-tasks/assesment_1/internal/model"
)

const (
	filePermission = 0o600
	dirPermission  = 0o750
)

func (uc *UseCase) Consume(ctx context.Context, errCh chan error) {
	outputDir := "./output"
	uc.l.InfoContext(ctx, "Checking if directory exists...", slog.String("outputDir", outputDir))

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		uc.l.InfoContext(ctx, "Directory does not exist. Creating...", slog.String("outputDir", outputDir))

		if err = os.MkdirAll(outputDir, dirPermission); err != nil {
			uc.l.InfoContext(ctx, "Failed to create directory", slog.String("outputDir", outputDir))
		}

		uc.l.InfoContext(ctx, "Directory created successfully", slog.String("outputDir", outputDir))
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			uc.l.InfoContext(ctx, "Shutting down...")
			uc.writeCacheToFile(errCh)
			return
		case <-ticker.C:
			uc.writeCacheToFile(errCh)
		}
	}
}

func (uc *UseCase) writeCacheToFile(errCh chan error) {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	for fileID, messages := range uc.cache {
		uc.writeMessagesToFile(fileID, messages, errCh)
		delete(uc.cache, fileID)
	}
}

// writeMessagesToFile записывает список сообщений в файл.
func (uc *UseCase) writeMessagesToFile(fileID string, messages []model.Message, errCh chan error) {
	f, err := os.OpenFile("output/"+fileID, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePermission)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", fileID, err)
		errCh <- err
	}

	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			fmt.Printf("Error closing file %s: %v\n", fileID, err)
		}
	}(f)

	for _, msg := range messages {
		if _, err = f.WriteString(msg.Data + "\n"); err != nil {
			fmt.Printf("Error writing to file %s: %v\n", fileID, err)
			errCh <- err
		}
	}

	fmt.Printf("Written %d messages to file %s\n", len(messages), fileID)
}
