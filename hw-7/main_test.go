package main

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"
)

// Тестовый HTTP-сервер для иммитации загрузки файлов
func createTestServer(t *testing.T, content string) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, content)
		assert.NoError(t, err)
	})
	return httptest.NewServer(handler)
}

// Тестирование downloadFile function
func TestDownloadFile(t *testing.T) {
	server := createTestServer(t, "test content")
	defer server.Close()

	task := DownloadTask{
		URL:      server.URL,
		FileName: "testfile.txt",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := downloadFile(ctx, task)
	assert.NoError(t, err)

	out, err := os.ReadFile(fmt.Sprintf("%s/%s", outputDir, task.FileName))
	assert.NoError(t, err)
	assert.Equal(t, "test content", string(out))

	// Cleanup
	err = os.Remove(fmt.Sprintf("%s/%s", outputDir, task.FileName))
	assert.NoError(t, err)
}

// Тестирование worker function
func TestWorker(t *testing.T) {
	server := createTestServer(t, "test content")
	defer server.Close()

	tasks := []DownloadTask{
		{URL: server.URL, FileName: "file1.txt"},
		{URL: server.URL, FileName: "file2.txt"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	taskChannel := make(chan DownloadTask, len(tasks))
	resultChannel := make(chan error, len(tasks))
	concurrencyLimiter := make(chan struct{}, concurrencyLimit)

	var wg sync.WaitGroup
	wg.Add(1)
	go worker(ctx, &wg, taskChannel, resultChannel, concurrencyLimiter)

	// Заполняем канал задачами
	for _, task := range tasks {
		taskChannel <- task
	}
	close(taskChannel)

	wg.Wait()
	close(resultChannel)

	for err := range resultChannel {
		assert.NoError(t, err)
	}

	for _, task := range tasks {
		out, err := os.ReadFile(fmt.Sprintf("%s/%s", outputDir, task.FileName))
		assert.NoError(t, err)
		assert.Equal(t, "test content", string(out))

		// Cleanup
		err = os.Remove(fmt.Sprintf("%s/%s", outputDir, task.FileName))
		assert.NoError(t, err)
	}
}

func TestMain(m *testing.M) {
	// Ensure output directory exists for tests
	err := os.MkdirAll(outputDir, dirPermissions)
	if err != nil {
		fmt.Printf("Failed to create output directory for tests: %v\n", err)
		os.Exit(1)
	}

	code := m.Run()

	// Cleanup output directory after tests
	err = os.RemoveAll(outputDir)
	if err != nil {
		fmt.Printf("Failed to remove output directory after tests: %v\n", err)
	}

	os.Exit(code)
}
