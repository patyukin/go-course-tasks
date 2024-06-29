package uploader

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/patyukin/go-course-tasks/hw-2/types"
)

func UploadHomeWorkData() types.InputData {
	file, err := os.Open("dz3.json")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Printf("Failed to close file: %s", err)
		}
	}(file)

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var data types.InputData

	// Парсинг JSON данных
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	return data
}
