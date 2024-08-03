package main

import (
	"testing"
)

/*
goos: darwin
goarch: arm64
pkg: github.com/patyukin/go-course-tasks/hw-8/ex-1
BenchmarkIsReflectMatrix
BenchmarkIsReflectMatrix-12                    	     381	   2833068 ns/op
BenchmarkIsReflectMatrixOptimized
BenchmarkIsReflectMatrixOptimized-12           	     549	   2117545 ns/op
BenchmarkIsReflectMatrixOConcurrent
BenchmarkIsReflectMatrixOConcurrent-12         	       5	 206448217 ns/op
BenchmarkIsReflectMatrixConcurrentMutex
BenchmarkIsReflectMatrixConcurrentMutex-12     	     615	   1960244 ns/op
BenchmarkIsReflectMatrixConcurrentAtomic
BenchmarkIsReflectMatrixConcurrentAtomic-12    	     594	   1967172 ns/op
PASS

Process finished with the exit code 0
*/

const n = 1000

func getMatrix() [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = i + j
			}
		}
	}

	return matrix
}

func BenchmarkIsReflectMatrix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsReflectMatrix(getMatrix())
	}
}
func BenchmarkIsReflectMatrixOptimized(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsReflectMatrixOptimized(getMatrix())
	}
}

func BenchmarkIsReflectMatrixOConcurrent(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsReflectMatrixConcurrency(getMatrix())
	}
}

func BenchmarkIsReflectMatrixConcurrentMutex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsReflectMatrixConcurrencyMutex(getMatrix())
	}
}

func BenchmarkIsReflectMatrixConcurrentAtomic(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsReflectMatrixConcurrencyAtomic(getMatrix())
	}
}
