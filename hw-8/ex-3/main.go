package main

import (
	"sync"
)

const numGoroutines = 10

func MinEl21(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	t1 := MinEl21(a[:len(a)/2])
	t2 := MinEl21(a[len(a)/2:])
	if t1 <= t2 {
		return t1
	}
	return t2
}

func MinEl22(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return minElHelper2(a, 0, len(a)-1)
}

func minElHelper2(a []int, start, end int) int {
	if start == end {
		return a[start]
	}
	mid := (start + end) / 2
	leftMin := minElHelper2(a, start, mid)
	rightMin := minElHelper2(a, mid+1, end)
	if leftMin <= rightMin {
		return leftMin
	}

	return rightMin
}

func findMin(arr []int, mu *sync.Mutex, minChan chan int) {
	minEl := arr[0]
	for _, v := range arr {
		if v < minEl {
			minEl = v
		}
	}

	mu.Lock()
	minChan <- minEl
	mu.Unlock()
}

// parallelMin эффективна, если длина слайса >= 10_000
func parallelMin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	length := len(arr)
	step := length / numGoroutines
	wg := &sync.WaitGroup{}
	minChan := make(chan int, numGoroutines)
	mu := &sync.Mutex{}

	for i := 0; i < numGoroutines; i++ {
		start := i * step
		end := start + step
		if i == numGoroutines-1 {
			end = length
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			findMin(arr[start:end], mu, minChan)
		}()
	}

	wg.Wait()
	close(minChan)

	globalMin := <-minChan
	for minEl := range minChan {
		if minEl < globalMin {
			globalMin = minEl
		}
	}

	return globalMin
}
