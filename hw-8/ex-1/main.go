package main

import (
	"sync"
	"sync/atomic"
)

func main() {

}

func IsReflectMatrix(a [][]int) bool {
	n := len(a)
	if n == 0 {
		return true
	}

	for i := 0; i < n-1; i++ {
		if len(a[i]) != n {
			return false
		}

		for j := 0; j < n; j++ {
			if a[i][j] != a[j][i] {
				return false
			}
		}
	}
	return true
}

func IsReflectMatrixOptimized(a [][]int) bool {
	n := len(a)
	if n == 0 {
		return true
	}

	for i := 0; i < n; i++ {
		if len(a[i]) != n {
			return false
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i][j] != a[j][i] {
				return false
			}
		}
	}
	return true
}

func IsReflectMatrixConcurrency(a [][]int) bool {
	n := len(a)
	if n == 0 {
		return true
	}

	for i := 0; i < n; i++ {
		if len(a[i]) != n {
			return false
		}
	}

	resultChan := make(chan bool, n*(n-1)/2)
	defer close(resultChan)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			go func(i, j int) {
				if a[i][j] != a[j][i] {
					resultChan <- false
					return
				}
				resultChan <- true
			}(i, j)
		}
	}

	for i := 0; i < n*(n-1)/2; i++ {
		if !<-resultChan {
			return false
		}
	}

	return true
}

func IsReflectMatrixConcurrencyMutex(a [][]int) bool {
	n := len(a)
	if n == 0 {
		return true
	}

	for i := 0; i < n; i++ {
		if len(a[i]) != n {
			return false
		}
	}

	wg := &sync.WaitGroup{}
	isSymmetric := true
	mu := &sync.Mutex{}

	checkSymmetry := func(start, end int) {
		defer wg.Done()
		for i := start; i < end && isSymmetric; i++ {
			for j := i + 1; j < n; j++ {
				if a[i][j] != a[j][i] {
					mu.Lock()
					isSymmetric = false
					mu.Unlock()
					return
				}
			}
		}
	}

	nWorkers := 8
	chunkSize := (n + nWorkers - 1) / nWorkers

	for i := 0; i < nWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > n {
			end = n
		}

		wg.Add(1)
		go checkSymmetry(start, end)
	}

	wg.Wait()
	return isSymmetric
}

func IsReflectMatrixConcurrencyAtomic(a [][]int) bool {
	n := len(a)
	if n == 0 {
		return true
	}

	for i := 0; i < n; i++ {
		if len(a[i]) != n {
			return false
		}
	}

	wg := &sync.WaitGroup{}
	var isSymmetric int32 = 1

	checkSymmetry := func(start, end int) {
		defer wg.Done()
		for i := start; i < end && atomic.LoadInt32(&isSymmetric) == 1; i++ {
			for j := i + 1; j < n; j++ {
				if a[i][j] != a[j][i] {
					atomic.StoreInt32(&isSymmetric, 0)
					return
				}
			}
		}
	}

	nWorkers := 8
	chunkSize := (n + nWorkers - 1) / nWorkers

	for i := 0; i < nWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > n {
			end = n
		}

		wg.Add(1)
		go checkSymmetry(start, end)
	}

	wg.Wait()
	return atomic.LoadInt32(&isSymmetric) == 1
}
