package main

import (
	"fmt"
	"io"
	"sync"
	"time"
)

type GoodLogger struct {
	output io.Writer
	pool   sync.Pool
}

// zerolog@v1.33.0
func NewGoodLogger(output io.Writer) *GoodLogger {
	return &GoodLogger{
		output: output,
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, 64)
			},
		},
	}
}

func (l *GoodLogger) Info(msg string) {
	buf := l.pool.Get().([]byte)
	buf = buf[:0]

	timestamp := time.Now().Format(time.RFC3339)
	buf = append(buf, fmt.Sprintf("[%s] %s: %s\n", timestamp, "INFO", msg)...)
	l.output.Write(buf)
	l.pool.Put(buf)
}
