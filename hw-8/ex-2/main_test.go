package main

import (
	"math/rand"
	"testing"
)

/*
goos: darwin
goarch: arm64
pkg: github.com/patyukin/go-course-tasks/hw-8/ex-2
BenchmarkMinElSmall
BenchmarkMinElSmall-12         	105401647	        11.26 ns/op
BenchmarkMinElMedium
BenchmarkMinElMedium-12        	 2759084	       385.2 ns/op
BenchmarkMinElLarge
BenchmarkMinElLarge-12         	   20330	     58813 ns/op
BenchmarkMinElLoopSmall
BenchmarkMinElLoopSmall-12     	288718388	         4.168 ns/op
BenchmarkMinElLoopMedium
BenchmarkMinElLoopMedium-12    	38959878	        30.67 ns/op
BenchmarkMinElLoopLarge
BenchmarkMinElLoopLarge-12     	  403670	      2929 ns/op
*/

// generateRandomSlice generates a slice of random integers of the given size.
func generateRandomSlice(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10000) // от 0 до 9999
	}
	return arr
}

func BenchmarkMinElSmall(b *testing.B) {
	arr := generateRandomSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl(arr)
	}
}

func BenchmarkMinElMedium(b *testing.B) {
	arr := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl(arr)
	}
}

func BenchmarkMinElLarge(b *testing.B) {
	arr := generateRandomSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl(arr)
	}
}

func BenchmarkMinElLoopSmall(b *testing.B) {
	arr := generateRandomSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinElRecLoop(arr)
	}
}

func BenchmarkMinElLoopMedium(b *testing.B) {
	arr := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinElRecLoop(arr)
	}
}

func BenchmarkMinElLoopLarge(b *testing.B) {
	arr := generateRandomSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinElRecLoop(arr)
	}
}
