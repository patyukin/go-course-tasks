package main

import (
	"fmt"
	"sync"
)

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	for _, ch := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func produce(x, y int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := x; i < x+y; i++ {
			ch <- i
		}

		close(ch)
	}()

	return ch
}

func main() {
	ch1 := produce(10, 3)
	ch2 := produce(20, 3)

	merged := merge(ch1, ch2)

	for val := range merged {
		fmt.Println(val)
	}
}
