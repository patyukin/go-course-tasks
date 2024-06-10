package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

const (
	permission = 0o644
	fileName   = "output.txt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dataCh := make(chan string)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		readInput(ctx, dataCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeToFile(ctx, dataCh)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("\nЗавершение работы программы...")
		cancel()
		wg.Wait()
	}()

	<-ctx.Done()

	fmt.Printf("\nПриложение завершено.\n")
}

func readInput(ctx context.Context, dataCh chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			var input string
			fmt.Print("Введите данные: ")
			_, err := fmt.Scanln(&input)
			if err != nil {
				log.Fatalf("Error reading input: %v", err)
			}

			dataCh <- input
		}
	}
}

func writeToFile(ctx context.Context, dataCh chan string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, permission)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(file)

	for {
		select {
		case <-ctx.Done():
			return
		case data := <-dataCh:
			if strings.TrimSpace(data) == "" {
				continue
			}

			if _, err = file.WriteString(data + "\n"); err != nil {
				log.Printf("Error writing to file: %v", err)
			}
		}
	}
}
