package main

import (
	"os"
	"testing"
	"time"
)

func TestValidateToken(t *testing.T) {
	valid := validateToken("token1")
	if !valid {
		t.Errorf("Expected token1 to be valid")
	}

	invalid := validateToken("invalid_token")
	if invalid {
		t.Errorf("Expected invalid_token to be invalid")
	}
}

func TestCacheMessage(t *testing.T) {
	msg := Message{Token: "token1", FileID: "file_test", Data: "Test data"}
	cacheMessage(msg)

	cache.mu.RLock()
	defer cache.mu.RUnlock()

	messages, exists := cache.messages["file_test"]
	if !exists {
		t.Errorf("Expected cache to contain messages for file_test")
	}

	if len(messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(messages))
	}

	if messages[0].Data != "Test data" {
		t.Errorf("Expected message data to be 'Test data', got %s", messages[0].Data)
	}
}

func TestWorkerWrite(t *testing.T) {
	msg := Message{Token: "token1", FileID: "file_test", Data: "Test data"}
	cacheMessage(msg)

	writeCacheToFiles()

	fileContent, err := os.ReadFile("file_test.txt")
	if err != nil {
		t.Fatalf("Expected file to be written but got error: %v", err)
	}

	expectedContent := "Test data\n"
	if string(fileContent) != expectedContent {
		t.Errorf("Expected file content to be '%s', got '%s'", expectedContent, string(fileContent))
	}

	cache.mu.RLock()
	defer cache.mu.RUnlock()

	if len(cache.messages["file_test"]) != 0 {
		t.Errorf("Expected cache to be cleared for file_test")
	}

	os.Remove("file_test.txt")
}

// Helper function to cache a message
func cacheMessage(msg Message) {
	if validateToken(msg.Token) {
		cache.mu.Lock()
		defer cache.mu.Unlock()
		cache.messages[msg.FileID] = append(cache.messages[msg.FileID], msg)
	}
}

func TestWorkerCycle(t *testing.T) {
	messageChannel := make(chan Message, 100)

	// Add messages to channel
	for i := 0; i < 3; i++ {
		msg := Message{Token: "token1", FileID: "file_test", Data: "Data " + string(rune(i+48))}
		messageChannel <- msg
	}

	// Start worker in a goroutine
	wg.Add(1)
	go worker(messageChannel)

	// Wait for worker to process messages
	time.Sleep(2 * time.Second)

	// Write remaining cache to files and shutdown worker
	close(messageChannel)
	wg.Wait()

	fileContent, err := os.ReadFile("file_test.txt")
	if err != nil {
		t.Fatalf("Expected file to be written but got error: %v", err)
	}

	expectedContent := "Data 0\nData 1\nData 2\n"
	if string(fileContent) != expectedContent {
		t.Errorf("Expected file content to be '%s', got '%s'", expectedContent, string(fileContent))
	}

	os.Remove("file_test.txt")
}
