package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	permission = 0o600
	fileName   = "input.txt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan string)

	go readInput(ctx, ch)
	go writeToFile(ch)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Printf("\nПриложение завершено.\n")
}

func readInput(ctx context.Context, ch chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение ввода...")
			close(ch)

			return
		default:
			if scanner.Scan() {
				ch <- scanner.Text()
			}
		}
	}
}

func writeToFile(ch <-chan string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, permission)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Ошибка закрытия файла:", err)
		}
	}(file)

	writer := bufio.NewWriter(file)
	defer func(writer *bufio.Writer) {
		err = writer.Flush()
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
		}
	}(writer)

	for text := range ch {
		_, err = writer.WriteString(text + "\n")
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}
	}
}
