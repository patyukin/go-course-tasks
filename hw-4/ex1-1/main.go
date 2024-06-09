package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	permission = 0o644
	fileName   = "output.txt"
)

func main() {
	messageChan := make(chan string)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go readInput(messageChan)
	go writeToFile(messageChan)

	<-signalChan
	fmt.Println("\nЗавершение работы программы...")
	close(messageChan)
	fmt.Printf("\nПриложение завершено.\n")
}

func readInput(ch chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Чтение из консоли:", text)
		ch <- text
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения с консоли:", err)
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
		fmt.Println("Запись в файл:", text)
	}
}
