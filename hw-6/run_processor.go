package main

import (
	"fmt"
	"sync"
)

func RunProcessor(wg *sync.WaitGroup, prices []map[string]float64, mu *sync.Mutex) {
	go func() {
		defer wg.Done()

		for _, price := range prices {
			mu.Lock()
			for key, value := range price {
				price[key] = value + 1
			}

			fmt.Println(price)
			mu.Unlock()
		}
	}()
}
