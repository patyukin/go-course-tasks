package main

import (
	"fmt"
	"time"
)

type Bad1Logger struct{}

func (l *Bad1Logger) Info(msg interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	var logMessage string
	switch v := msg.(type) {
	case string:
		logMessage = fmt.Sprintf("[%s] %s: %s\n", timestamp, "INFO", v)
	default:
		logMessage = fmt.Sprintf("[%s] %s: %s\n", timestamp, "INFO", v)
	}

	fmt.Print(logMessage)
}
