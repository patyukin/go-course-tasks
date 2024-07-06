package main

import (
	"os"
	"testing"
)

func BenchmarkBad1Logger(b *testing.B) {
	logger := Bad1Logger{}
	for i := 0; i < b.N; i++ {
		logger.Info("Benchmark test message")
	}
}

func BenchmarkGood1Logger(b *testing.B) {
	var logger = NewGoodLogger(os.Stdout)
	for i := 0; i < b.N; i++ {
		logger.Info("Benchmark test message")
	}
}
