package main

import (
	"maps"
	"time"
)

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
			for key, value := range currentPrice {
				currentPrice[key] = value + 1
			}

			prices <- maps.Clone(currentPrice)
			time.Sleep(time.Second)
		}

		close(prices)
	}()

	return prices
}
