package main

import (
	"fmt"
	"sync"
)

const (
	processorCount = 3
)

func main() {
	p := RunWriter()
	prices := make([]map[string]float64, 0, processorCount)

	for price := range p {
		prices = append(prices, price)
	}

	for _, price := range prices {
		fmt.Println(price)
	}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(3)
	RunProcessor(wg, prices, mu)
	RunProcessor(wg, prices, mu)
	RunProcessor(wg, prices, mu)

	wg.Wait()
}
