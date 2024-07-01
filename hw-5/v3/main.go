package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	permission = 0o644
	fileName   = "output.txt"
)

// выход только через Enter
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		readInput(ctx, ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeToFile(ctx, ch)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	fmt.Println("\nЗавершение работы программы...")

	cancel()
	wg.Wait()

	fmt.Printf("\nПриложение завершено.\n")
}

func readInput(ctx context.Context, ch chan<- string) {
	defer close(ch)
	reader := bufio.NewReader(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Чтение ввода прервано")
			return
		default:
		}

		fmt.Print("Введите данные: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}

		ch <- line
	}
}

func writeToFile(ctx context.Context, ch <-chan string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, permission)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Printf("Failed to close file: %v", err)
		}
	}(file)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Чтение ввода прервано")
			return
		default:
		}

		line := <-ch
		if _, err = file.WriteString(line); err != nil {
			log.Fatal(err)
		}
	}
}
