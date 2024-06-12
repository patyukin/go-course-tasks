package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	permission = 0o644
	fileName   = "output.txt"
)

func main() {
	ch := make(chan string)

	go readInput(ch)
	go writeToFile(ch)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	fmt.Println("\nЗавершение работы программы...")
	fmt.Printf("\nПриложение завершено.\n")
}

func readInput(ch chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	defer close(ch)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}

		ch <- line
	}
}

func writeToFile(ch <-chan string) {
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
		line := <-ch
		if _, err = file.WriteString(line); err != nil {
			log.Fatal(err)
		}
	}
}
