package main

import (
	"sync"
	"testing"
)

func TestRunProcessor(t *testing.T) {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(1)

	prices := []map[string]float64{
		{"inst1": 1.1, "inst2": 2.1},
		{"inst3": 3.1, "inst4": 4.1},
	}

	expectedPrices := []map[string]float64{
		{"inst1": 2.1, "inst2": 3.1},
		{"inst3": 4.1, "inst4": 5.1},
	}

	RunProcessor(wg, prices, mu)

	wg.Wait()

	for i, price := range prices {
		for key, value := range price {
			if expectedPrices[i][key] != value {
				t.Errorf("Expected %f but got %f for key %s at index %d", expectedPrices[i][key], value, key, i)
			}
		}
	}
}
