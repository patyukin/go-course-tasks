package main

import (
	"os"
)

func main() {
	l1 := Bad1Logger{}
	l1.Info("Test message")

	l2 := NewGoodLogger(os.Stdout)
	l2.Info("Test message")
}
