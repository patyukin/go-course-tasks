package main

import (
	"math/rand"
	"testing"
)

/*
prev
BenchmarkMinElSmall
BenchmarkMinElSmall-12         		123068041	    10.33 ns/op
BenchmarkMinElMedium
BenchmarkMinElMedium-12        	 	2514009	       	426.7 ns/op
BenchmarkMinElLarge
BenchmarkMinElLarge-12         	  	204958	      	5766 ns/op
BenchmarkMinElLoopSmall
BenchmarkMinElLoopSmall-12     		284084210	    4.156 ns/op
BenchmarkMinElLoopMedium
BenchmarkMinElLoopMedium-12    		39631754	    30.64 ns/op
BenchmarkMinElLoopLarge
BenchmarkMinElLoopLarge-12     		3960752	       	298.1 ns/op

current
goos: darwin
goarch: arm64
pkg: github.com/patyukin/go-course-tasks/hw-8/ex-3
BenchmarkMinElSmall
BenchmarkMinElSmall-12                 	42723294	        24.80 ns/op
BenchmarkMinElMedium
BenchmarkMinElMedium-12                	 4358715	       277.5 ns/op
BenchmarkMinElLarge
BenchmarkMinElLarge-12                 	   27804	     40851 ns/op
BenchmarkMinEl2Small
BenchmarkMinEl2Small-12                	49314140	        23.05 ns/op
BenchmarkMinEl2Medium
BenchmarkMinEl2Medium-12               	 5128497	       233.1 ns/op
BenchmarkMinEl2Large
BenchmarkMinEl2Large-12                	   28819	     39641 ns/op
BenchmarkMinEl2ConcurrencySmall
BenchmarkMinEl2ConcurrencySmall-12     	  571381	      1917 ns/op
BenchmarkMinEl2ConcurrencyMedium
BenchmarkMinEl2ConcurrencyMedium-12    	  604867	      1927 ns/op
BenchmarkMinEl2LConcurrencyarge
BenchmarkMinEl2LConcurrencyarge-12     	  163738	      7442 ns/op
PASS

Process finished with the exit code 0
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
		MinEl21(arr)
	}
}

func BenchmarkMinElMedium(b *testing.B) {
	arr := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl21(arr)
	}
}

func BenchmarkMinElLarge(b *testing.B) {
	arr := generateRandomSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl21(arr)
	}
}

func BenchmarkMinEl2Small(b *testing.B) {
	arr := generateRandomSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl22(arr)
	}
}

func BenchmarkMinEl2Medium(b *testing.B) {
	arr := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl22(arr)
	}
}

func BenchmarkMinEl2Large(b *testing.B) {
	arr := generateRandomSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl22(arr)
	}
}

func BenchmarkMinEl2ConcurrencySmall(b *testing.B) {
	arr := generateRandomSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parallelMin(arr)
	}
}

func BenchmarkMinEl2ConcurrencyMedium(b *testing.B) {
	arr := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parallelMin(arr)
	}
}

func BenchmarkMinEl2LConcurrencyarge(b *testing.B) {
	arr := generateRandomSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parallelMin(arr)
	}
}
