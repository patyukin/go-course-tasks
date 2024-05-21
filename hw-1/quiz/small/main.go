package main

import (
	"crypto/rand"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

const (
	countColumns = 2
)

// Функция для чтения CSV файла.
func readCSV(filename string) ([][]string, error) {
	if filename == "" {
		filename = "problems.csv"
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии CSV файла: %w", err)
	}

	defer func(file *os.File) {
		closeError := file.Close()
		if closeError != nil {
			fmt.Printf("ошибка при закрытии CSV файла: %s\n", closeError)
		}
	}(file)

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении CSV файла: %w", err)
	}

	return records, nil
}

func shuffle(data [][]string) error {
	n := len(data)
	for idx := range n {
		j, err := cryptoRandInt(n)
		if err != nil {
			return fmt.Errorf("ошибка: %w", err)
		}

		data[idx], data[j] = data[j], data[idx]
	}

	return nil
}

func cryptoRandInt(max int) (int, error) {
	bigN, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("ошибка: %w", err)
	}

	return int(bigN.Int64()), nil
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "Csv файл в формате 'question,answer'")
	shuffleQuestions := flag.Bool("shuffle", false, "Перемешать вопросы?")
	flag.Parse()

	// Читаем вопросы и ответы из CSV файла
	records, err := readCSV(*csvFilename)
	if err != nil {
		log.Fatalf("Ошибка при открытии CSV файла: %s\n", *csvFilename)
	}

	// Перемешиваем вопросы, если это указано в флагах
	if *shuffleQuestions {
		shuffleError := shuffle(records)
		if shuffleError != nil {
			log.Fatalf("Невозможно перемешать вопросы: %s\n", shuffleError)
		}
	}

	correct := 0
	incorrect := 0

	for idx, record := range records {
		if len(record) != countColumns {
			fmt.Printf("Неверный формат записи, пропуск: %s", record)
			continue
		}

		question := record[0]
		correctAnswer := record[1]

		var userAnswer string

		fmt.Printf("Вопрос #%d: %s = ", idx+1, question)

		_, scanError := fmt.Scanf("%s\n", &userAnswer)
		if scanError != nil {
			log.Fatalf("Ошибка чтения ответа: %s\n", scanError)
		}

		if strings.EqualFold(strings.TrimSpace(userAnswer), strings.TrimSpace(correctAnswer)) {
			correct++
			continue
		}

		incorrect++
	}

	fmt.Printf("Правильных ответов: %d\n", correct)
	fmt.Printf("Неправильных ответов: %d\n", incorrect)
}
