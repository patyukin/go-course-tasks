package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Message struct {
	Token  string
	FileID string
	Data   string
}

var validTokens = map[string]bool{
	"token1": true,
	"token2": true,
	"token3": true,
}

type Cache struct {
	mu       sync.RWMutex
	messages map[string][]Message
}

var cache = Cache{
	mu:       sync.RWMutex{},
	messages: make(map[string][]Message),
}

var wg = &sync.WaitGroup{}

func main() {
	ch := make(chan Message, 100500)

	go simulateUsers(ch)

	wg.Add(1)
	go worker(ch)

	setupGracefulShutdown(ch)

	wg.Wait()
}

func simulateUsers(ch chan<- Message) {
	users := []string{"user1", "user2", "user3"}
	tokens := []string{"token1", "token2", "invalid_token"}

	for {
		for i := 0; i < len(users); i++ {
			ch <- Message{
				Token:  tokens[i],
				FileID: "file_" + fmt.Sprint(i),
				Data:   "data from " + users[i],
			}

			time.Sleep(200 * time.Millisecond)
		}
	}
}

func worker(ch <-chan Message) {
	defer wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case msg := <-ch:
			if validateToken(msg.Token) {
				cache.mu.RLock()
				cache.messages[msg.FileID] = append(cache.messages[msg.FileID], msg)
				cache.mu.RUnlock()
			}
		case <-ticker.C:
			writeCacheToFiles()
		}
	}
}

func validateToken(token string) bool {
	valid, exists := validTokens[token]
	return exists && valid
}

func writeCacheToFiles() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for fileID, messages := range cache.messages {
		writeMessagesToFile(fileID, messages)
	}

	cache.messages = make(map[string][]Message)
}

func writeMessagesToFile(fileID string, messages []Message) {
	file, err := os.OpenFile(fileID+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for _, msg := range messages {
		_, err = file.WriteString(msg.Data + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Printf("Wrote %d messages to %s.txt\n", len(messages), fileID)
}

func setupGracefulShutdown(ch chan<- Message) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("Graceful shutdown initiated...")
		close(ch)

		// Write remaining cache to files
		writeCacheToFiles()

		wg.Done()
	}()
}
