package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	squares := make([]int, 0, len(nums))
	cubes := make([]int, 0, len(nums))

	wg.Add(1)
	go func() {
		defer wg.Done()

		ch := calcSquares(nums)
		for num := range ch {
			squares = append(squares, num)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		ch := calcCubes(nums)
		for num := range ch {
			cubes = append(cubes, num)
		}
	}()

	wg.Wait()
	fmt.Printf("squares: %v\n", squares)
	fmt.Printf("cubes: %v\n", cubes)
}

func calcCubes(nums []int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, num := range nums {
			ch <- num * num * num
		}
	}()

	return ch
}

func calcSquares(nums []int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, num := range nums {
			ch <- num * num
		}
	}()

	return ch
}
