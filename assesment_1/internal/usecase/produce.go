package usecase

import (
	"context"
	"fmt"
	"github.com/patyukin/go-course-tasks/assesment_1/internal/model"
	"log/slog"
)

func (uc *UseCase) Produce(ctx context.Context, errCh chan error) {
	for _, ch := range uc.inputChs {
		go uc.dispatcher(ctx, ch, errCh)
	}

	for i, message := range uc.messages {
		uc.sendMessage(i%uc.cfg.InputChsCount, message.Token, message.FileID, message.Data)
	}
}

func (uc *UseCase) sendMessage(channelIndex int, token, fileID, data string) {
	if channelIndex < 0 || channelIndex >= len(uc.inputChs) {
		fmt.Printf("Invalid channel index: %d\n", channelIndex)
		return
	}

	uc.inputChs[channelIndex] <- model.Message{Token: token, FileID: fileID, Data: data}
}

func (uc *UseCase) closeInputChs() {
	for _, ch := range uc.inputChs {
		close(ch)
	}
}

func (uc *UseCase) dispatcher(ctx context.Context, ch chan model.Message, errCh chan error) {
	for {
		select {
		case <-ctx.Done():
			uc.l.InfoContext(ctx, "Shutting down...", slog.String("from", "Produce => dispatcher"))
			uc.closeInputChs()
			return
		case msg := <-ch:
			uc.cacheMessage(msg, errCh)
		}
	}
}

func (uc *UseCase) cacheMessage(msg model.Message, errCh chan error) {
	if !uc.validateToken(msg.Token) {
		fmt.Printf("Invalid token: %s\n", msg.Token)
		return
	}

	uc.mu.Lock()
	uc.cache[msg.FileID] = append(uc.cache[msg.FileID], msg)
	uc.mu.Unlock()
}

// validateToken проверяет валидность токена
func (uc *UseCase) validateToken(token string) bool {
	_, ok := uc.validTokens[token]
	return ok
}
