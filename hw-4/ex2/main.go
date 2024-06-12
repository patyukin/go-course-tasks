package main

import (
	"fmt"
	"sync"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func separateNumbers(numbers []int, primeCh, compositeCh chan<- int) {
	for _, num := range numbers {
		if isPrime(num) {
			primeCh <- num
		} else {
			compositeCh <- num
		}
	}

	close(primeCh)
	close(compositeCh)
}

func main() {
	numbers := []int{2, 4, 6, 7, 8, 11, 13, 15, 17, 18, 19, 20}
	var primeNumbers, compositeNumbers []int

	primeCh := make(chan int)
	compositeCh := make(chan int)

	go separateNumbers(numbers, primeCh, compositeCh)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range primeCh {
			primeNumbers = append(primeNumbers, num)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range compositeCh {
			compositeNumbers = append(compositeNumbers, num)
		}
	}()

	wg.Wait()

	fmt.Println("Простые числа:", primeNumbers)
	fmt.Println("Составные числа:", compositeNumbers)
}
