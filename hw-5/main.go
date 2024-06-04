package main

import (
	"fmt"
	"sync"
)

func RunProcessor(wg *sync.WaitGroup, prices []map[string]float64, mu *sync.Mutex) {
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		for _, price := range prices {
			for key, value := range price {
				price[key] = value + 1
			}
			fmt.Println(price)
		}
	}()
}

func RunWriter() <-chan map[string]float64 {
	prices := make(chan map[string]float64)
	go func() {
		currentPrice := map[string]float64{
			"inst1": 1.1,
			"inst2": 2.1,
			"inst3": 3.1,
			"inst4": 4.1,
		}

		for i := 1; i < 5; i++ {
			temp := make(map[string]float64)
			for key, value := range currentPrice {
				temp[key] = value + 1
			}
			prices <- temp
		}
		close(prices)
	}()
	return prices
}

func main() {
	p := RunWriter()
	prices := make([]map[string]float64, 0)

	for price := range p {
		prices = append(prices, price)
	}

	for _, price := range prices {
		fmt.Println(price)
	}

	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	RunProcessor(wg, prices, mu)
	RunProcessor(wg, prices, mu)
	RunProcessor(wg, prices, mu)
	wg.Wait()
}
