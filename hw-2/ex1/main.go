package main

import (
	"fmt"
	"sort"
)

func intersect(slices ...[]int) []int {
	if len(slices) == 0 {
		return []int{}
	}

	counts := make(map[int]int)

	for _, slice := range slices {
		if len(slice) == 0 {
			return []int{}
		}

		unique := make(map[int]bool)

		for _, num := range slice {
			unique[num] = true
		}

		for num := range unique {
			counts[num]++
		}
	}

	var result []int

	for num, count := range counts {
		if count == len(slices) {
			result = append(result, num)
		}
	}

	sort.Ints(result)

	return result
}

func main() {
	fmt.Print(intersect([]int{1, 2, 3, 2}, []int{3, 2}, []int{1, 2, 3, 2}))
}
