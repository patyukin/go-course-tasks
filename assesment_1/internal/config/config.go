package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Env           string
	InputChsCount int
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	cfg := &Config{}

	cfg.Env = os.Getenv("ENV")
	inputChsCount, err := strconv.Atoi(os.Getenv("INPUT_CHANNELS_COUNT"))
	if err != nil {
		return nil, fmt.Errorf("error parsing INPUT_CHANNELS_COUNT: %v", err)
	}

	cfg.InputChsCount = inputChsCount

	return cfg, nil
}
