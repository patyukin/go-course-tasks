package main

import (
	"fmt"
)

// Numbers — слайс численных типов.
type Numbers[T int | float64 | float32] []T

// Sum возвращает сумму всех элементов в слайсе.
func (n *Numbers[T]) Sum() T {
	if n == nil {
		return 0
	}

	var sum T
	for _, v := range *n {
		sum += v
	}

	return sum
}

// Product возвращает произведение всех элементов в слайсе.
func (n *Numbers[T]) Product() T {
	if n == nil {
		return 0
	}

	var product T = 1
	for _, v := range *n {
		product *= v
	}

	return product
}

// Equals сравнивает c другим слайсом на равенство.
func (n *Numbers[T]) Equals(other *Numbers[T]) bool {
	if n == nil && other == nil {
		return true
	}

	if n == nil || other == nil {
		return false
	}

	if len(*n) != len(*other) {
		return false
	}

	for i, v := range *n {
		if v != (*other)[i] {
			return false
		}
	}

	return true
}

// IndexOf проверяет, является ли аргумент элементом слайса, и если да - возвращает индекс первого найденного элемента.
func (n *Numbers[T]) IndexOf(value T) (int, bool) {
	for i, v := range *n {
		if v == value {
			return i, true
		}
	}

	return -1, false
}

// RemoveByValue удаляет элемент слайса по значению.
func (n *Numbers[T]) RemoveByValue(value T) bool {
	index, found := n.IndexOf(value)
	if !found {
		return false
	}
	return n.RemoveByIndex(index)
}

// RemoveByIndex удаляет элемент слайса по индексу.
func (n *Numbers[T]) RemoveByIndex(index int) bool {
	if index < 0 || index >= len(*n) {
		return false
	}

	*n = append((*n)[:index], (*n)[index+1:]...)
	return true
}

func main() {
	nums := Numbers[int]{1, 2, 3, 4, 5}

	fmt.Println("Sum:", nums.Sum())

	fmt.Println("Product:", nums.Product())

	otherNums := Numbers[int]{1, 2, 3, 4, 5}
	fmt.Println("Equals:", nums.Equals(&otherNums))

	otherNums = Numbers[int]{1, 2, 3, 4}
	fmt.Println("Equals:", nums.Equals(&otherNums))

	if idx, found := nums.IndexOf(3); found {
		fmt.Printf("IndexOf 3: %d\n", idx)
	} else {
		fmt.Println("3 not found")
	}

	_ = nums.RemoveByValue(3)
	fmt.Println("After removing value 3:", nums)

	_ = nums.RemoveByIndex(2)
	fmt.Println("After removing index 2:", nums)
}
