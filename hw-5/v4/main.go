package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	permission = 0o644
	fileName   = "output.txt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	ch := make(chan string)

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

	go func() {
		cancel()
		wg.Wait()
	}()

	fmt.Printf("\nПриложение завершено.\n")
}

func readInput(ctx context.Context, ch chan<- string) {
	defer close(ch)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Чтение ввода прервано")
			return
		default:
		}

		fmt.Print("Введите данные: ")
		if scanner.Scan() {
			input := scanner.Text()
			ch <- input
		}
	}
}

func writeToFile(ctx context.Context, inputChan <-chan string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, permission)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}

	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			fmt.Println("Ошибка при закрытии файла:", err)
		}
	}(file)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Запись в файл прервана")
			return
		case input, ok := <-inputChan:
			if !ok {
				return
			}

			if _, err = file.WriteString(input + "\n"); err != nil {
				fmt.Println("Ошибка записи в файл:", err)
				return
			}
		}
	}
}
