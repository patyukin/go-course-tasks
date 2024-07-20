package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	countColumns = 2
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите название CSV файла: ")
	scanner.Scan()
	filename := scanner.Text()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Ошибка при закрытии файла: %s\n", err)
		}
	}(file)

	// Создаем буферизованный ридер
	bufferedReader := bufio.NewReader(file)

	// Создаем ридер CSV поверх буферизованного ридера
	reader := csv.NewReader(bufferedReader)

	correct := 0
	incorrect := 0

	for {
		record, readError := reader.Read()
		if readError != nil {
			if readError == io.EOF {
				break
			}

			fmt.Println("Ошибка при чтении файла:", readError)

			return
		}

		if len(record) != countColumns {
			fmt.Printf("Неверный формат записи, пропуск: %s", record)

			continue
		}

		question := record[0]
		correctAnswer := record[1]

		fmt.Printf("%s: ", question)
		scanner.Scan()
		userAnswer := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(userAnswer), strings.TrimSpace(correctAnswer)) {
			correct++
			continue
		}

		incorrect++
	}

	fmt.Printf("Правильных ответов: %d\n", correct)
	fmt.Printf("Неправильных ответов: %d\n", incorrect)
}
