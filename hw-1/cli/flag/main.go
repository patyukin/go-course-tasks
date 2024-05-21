package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Определение флагов командной строки
	createFlag := flag.String("create", "", "Создать файл: укажите имя файла")
	readFlag := flag.String("read", "", "Прочитать содержимое файла: укажите имя файла")
	deleteFlag := flag.String("delete", "", "Удалить файл: укажите имя файла")

	flag.Parse()

	if *createFlag != "" {
		err := os.WriteFile(*createFlag, []byte("Пример содержимого файла"), 0644)
		if err != nil {
			fmt.Println("Ошибка при создании файла:", err)
		} else {
			fmt.Println("Файл успешно создан:", *createFlag)
		}
	}

	if *readFlag != "" {
		data, err := os.ReadFile(*readFlag)
		if err != nil {
			fmt.Println("Ошибка при чтении файла:", err)
		} else {
			fmt.Println("Содержимое файла", *readFlag, ":\n", string(data))
		}
	}

	if *deleteFlag != "" {
		err := os.Remove(*deleteFlag)
		if err != nil {
			fmt.Println("Ошибка при удалении файла:", err)
		} else {
			fmt.Println("Файл успешно удален:", *deleteFlag)
		}
	}
}
