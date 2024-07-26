package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	workerCount      = 20
	concurrencyLimit = 4
	outputDir        = "./downloads"
	dirPermissions   = 0o750
)

// DownloadTask represents a file download task.
type DownloadTask struct {
	URL      string
	FileName string
}

// downloadFile performs the file download.
func downloadFile(ctx context.Context, task DownloadTask) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, task.URL, http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	time.Sleep(1 * time.Second)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", task.URL, err)
	}

	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			fmt.Printf("Failed to close response body: %v\n", err)
		}
	}(resp.Body)

	if ctx.Err() != nil {
		return fmt.Errorf("failed to download %s: %w", task.URL, ctx.Err())
	}

	// Create output file
	out, err := os.Create(fmt.Sprintf("%s/%s", outputDir, task.FileName))
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", task.FileName, err)
	}

	defer func(out *os.File) {
		if err = out.Close(); err != nil {
			fmt.Printf("Failed to close output file: %v\n", err)
		}
	}(out)

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write %v to file %s: %w", task.URL, task.FileName, err)
	}

	return nil
}

// worker function processes download tasks from the input channel and sends results to the output channel.
func worker(ctx context.Context, wg *sync.WaitGroup, tasks <-chan DownloadTask, results chan<- error,
	concurrencyLimiter chan struct{},
) {
	defer wg.Done()

	for task := range tasks {
		select {
		case <-ctx.Done():
			results <- ctx.Err()
			return
		default:
		}

		// limit concurrency
		concurrencyLimiter <- struct{}{}
		err := downloadFile(ctx, task)
		results <- err
		//  after completing task
		<-concurrencyLimiter
	}
}

func main() {
	// Ensure output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.Mkdir(outputDir, dirPermissions)
		if err != nil {
			fmt.Printf("Failed to create output directory: %v\n", err)
			return
		}
	}

	// List of download tasks
	tasks := []DownloadTask{
		{URL: "https://example.com/file1.jpg", FileName: "file1.jpeg"},
		{URL: "https://example.com/file2.jpg", FileName: "file2.jpeg"},
		{URL: "https://example.com/file3.jpg", FileName: "file3.jpeg"},
		{URL: "https://example.com/file4.jpg", FileName: "file4.jpeg"},
		{URL: "https://example.com/file5.jpg", FileName: "file5.jpeg"},
		{URL: "https://example.com/file6.jpg", FileName: "file6.jpeg"},
		{URL: "https://example.com/file7.jpg", FileName: "file7.jpeg"},
		{URL: "https://example.com/file8.jpg", FileName: "file8.jpeg"},
		{URL: "https://example.com/file9.jpg", FileName: "file9.jpeg"},
		{URL: "https://example.com/file10.jpg", FileName: "file10.jpeg"},
		{URL: "https://example.com/file11.jpg", FileName: "file11.jpeg"},
		{URL: "https://example.com/file12.jpg", FileName: "file12.jpeg"},
		{URL: "https://example.com/file13.jpg", FileName: "file13.jpeg"},
		{URL: "https://example.com/file14.jpg", FileName: "file14.jpeg"},
		{URL: "https://example.com/file15.jpg", FileName: "file15.jpeg"},
		{URL: "https://example.com/file16.jpg", FileName: "file16.jpeg"},
		{URL: "https://example.com/file17.jpg", FileName: "file17.jpeg"},
		{URL: "https://example.com/file18.jpg", FileName: "file18.jpeg"},
		{URL: "https://example.com/file19.jpg", FileName: "file19.jpeg"},
		{URL: "https://example.com/file20.jpg", FileName: "file20.jpeg"},
	}

	taskChannel := make(chan DownloadTask, len(tasks))
	resultChannel := make(chan error, len(tasks))

	// Start the timer
	startTime := time.Now()

	concurrencyLimiter := make(chan struct{}, concurrencyLimit)

	// Worker timeout duration starts here
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	wg := &sync.WaitGroup{}
	for range workerCount {
		wg.Add(1)
		go worker(ctx, wg, taskChannel, resultChannel, concurrencyLimiter)
	}

	// Send tasks to workers
	for _, task := range tasks {
		taskChannel <- task
	}
	close(taskChannel)

	// Wait for all workers to finish
	wg.Wait()
	close(resultChannel)

	// Check results
	for err := range resultChannel {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	// End the timer
	elapsedTime := time.Since(startTime)
	fmt.Printf("All downloads completed in %s\n", elapsedTime)
}
