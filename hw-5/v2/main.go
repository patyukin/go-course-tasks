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

func readInput(ctx context.Context, ch chan string) {
	defer close(ch)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите данные: ")
		tempCh := make(chan string, 1)

		go func() {
			input, _ := reader.ReadString('\n')
			tempCh <- input
		}()

		select {
		case <-ctx.Done():
			fmt.Println("Чтение ввода прервано")
			return
		case input := <-tempCh:
			ch <- input
			if input == "" {
				return
			}
		}
	}
}

func writeToFile(ch chan string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, permission)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Ошибка при закрытии файла:", err)
		}
	}(file)

	for input := range ch {
		if _, err = file.WriteString(input + "\n"); err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}
	}
}

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
		writeToFile(ch, fileName)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	fmt.Println("\nЗавершение работы программы...")

	cancel()
	wg.Wait()
	fmt.Printf("\nПриложение завершено.\n")
}
