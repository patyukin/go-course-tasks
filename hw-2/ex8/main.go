package main

func IsEqualArrays[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	counts := make(map[T]int)
	for _, v := range slice1 {
		counts[v]++
	}

	for _, v := range slice2 {
		counts[v]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{3, 4, 2, 1}
	arr3 := []int{1, 2, 3}
	arr4 := []int{1, 2, 3, 3, 4}

	println(IsEqualArrays(arr1, arr2))
	println(IsEqualArrays(arr1, arr3))
	println(IsEqualArrays(arr1, arr4))
}
