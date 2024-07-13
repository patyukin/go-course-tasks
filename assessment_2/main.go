package main

import (
	"fmt"
)

func EvalSequence(matrix [][]int, userAnswer []int) int {
	// validation
	if !validateInput(matrix, userAnswer) {
		return 0
	}

	maxGrade := calcMaxGrade(matrix)
	userGrade := calcUserGrade(matrix, userAnswer)

	if maxGrade == 0 {
		return 0
	}

	percent := userGrade * 100 / maxGrade
	return percent
}

// dfsForMaxGrade - dfs для поиска максимального балла
func dfsForMaxGrade(matrix [][]int, node int, currentGrade int, visited []bool, maxGrade *int) {
	n := len(matrix)
	visited[node] = true
	isEnd := true

	for i := 0; i < n; i++ {
		if !visited[i] && matrix[node][i] > 0 {
			isEnd = false
			dfsForMaxGrade(matrix, i, currentGrade+matrix[node][i], visited, maxGrade)
		}
	}

	if isEnd && currentGrade > *maxGrade {
		*maxGrade = currentGrade
	}
	visited[node] = false
}

func dfsForMaxGradeIterative(matrix [][]int, node int, maxGrade *int) {
	n := len(matrix)
	stack := []struct {
		node         int
		currentGrade int
		visited      []bool
	}{{node, 0, make([]bool, n)}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		current.visited[current.node] = true
		hasUnvisitedNeighbor := false

		for i := 0; i < n; i++ {
			if !current.visited[i] && matrix[current.node][i] > 0 {
				hasUnvisitedNeighbor = true
				newVisited := make([]bool, n)
				copy(newVisited, current.visited)
				stack = append(stack, struct {
					node         int
					currentGrade int
					visited      []bool
				}{i, current.currentGrade + matrix[current.node][i], newVisited})
			}
		}

		if !hasUnvisitedNeighbor {
			if current.currentGrade > *maxGrade {
				*maxGrade = current.currentGrade
			}
		}
	}
}

// dfsForMaxGrade - dfs для поиска максимального балла
func dfsForMaxGradeIterativeMemo(matrix [][]int, node int, maxGrade *int, memo map[int]int) {
	n := len(matrix)
	stack := []struct {
		node         int
		currentGrade int
		visited      []bool
	}{{node, 0, make([]bool, n)}}

	for len(stack) > 0 {
		// Извлечение текущего состояния из стека
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Проверка кэшированного значения
		if cachedGrade, exists := memo[current.node]; exists {
			if current.currentGrade+cachedGrade > *maxGrade {
				*maxGrade = current.currentGrade + cachedGrade
			}
			continue
		}

		current.visited[current.node] = true
		hasUnvisitedNeighbor := false
		for i := 0; i < n; i++ {
			if !current.visited[i] && matrix[current.node][i] > 0 {
				hasUnvisitedNeighbor = true
				newVisited := make([]bool, n)
				copy(newVisited, current.visited)
				stack = append(stack, struct {
					node         int
					currentGrade int
					visited      []bool
				}{i, current.currentGrade + matrix[current.node][i], newVisited})
			}
		}

		if !hasUnvisitedNeighbor {
			if current.currentGrade > *maxGrade {
				*maxGrade = current.currentGrade
			}
			memo[current.node] = current.currentGrade
		}
	}
}

// calcMaxGrade - расчет максимального балла
func calcMaxGrade(matrix [][]int) int {
	maxGrade := 0
	n := len(matrix)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dfsForMaxGrade(matrix, i, 0, visited, &maxGrade)
	}

	return maxGrade
}

// calcUserGrade - расчет баллов пользователя
func calcUserGrade(matrix [][]int, userAnswer []int) int {
	userGrade := 0
	for i := 0; i < len(userAnswer)-1; i++ {
		from := userAnswer[i]
		to := userAnswer[i+1]
		userGrade += matrix[from][to]
	}

	return userGrade
}

// validateInput - валидация входных данных
func validateInput(matrix [][]int, userAnswer []int) bool {
	size := len(matrix)
	nonZeroFound := false

	if size == 0 {
		fmt.Println("пустая матрица")
		return false
	}

	for i := 0; i < size; i++ {
		if len(matrix[i]) != size {
			fmt.Println("матрица не квадратная")
			return false
		}

		if matrix[i][i] != 0 {
			fmt.Println("матрица содержит петли (ненулевая диагональ)")
			return false
		}

		for j := i + 1; j < size; j++ {
			if matrix[i][j] != matrix[j][i] {
				fmt.Println("матрица не симметрична")
				return false
			}

			if matrix[i][j] < 0 {
				fmt.Println("матрица содержит отрицательные значения")
				return false
			}

			if matrix[i][j] != 0 {
				nonZeroFound = true
			}
		}
	}

	if !nonZeroFound {
		fmt.Println("матрица содержит только нулевые значения")
		return false
	}

	seen := make(map[int]struct{})
	for _, answer := range userAnswer {
		if answer < 0 || answer >= size {
			fmt.Println("ответы пользователя выходят за пределы допустимого диапазона")
			return false
		}

		if _, ok := seen[answer]; ok {
			fmt.Println("повторяющиеся ответы")
			return false
		}

		seen[answer] = struct{}{}
	}

	// Дополнительная проверка на связность графа
	if !isGraphConnected(matrix) {
		fmt.Println("граф не связан")
		return false
	}

	return true
}

func performDFS(matrix [][]int, node int, visited []bool) {
	visited[node] = true
	size := len(matrix)
	for i := 0; i < size; i++ {
		if matrix[node][i] > 0 && !visited[i] {
			performDFS(matrix, i, visited)
		}
	}
}

// isGraphConnected - проверка связности графа
func isGraphConnected(matrix [][]int) bool {
	size := len(matrix)
	visited := make([]bool, size)

	performDFS(matrix, 0, visited)

	for i := 0; i < size; i++ {
		if !visited[i] {
			return false
		}
	}

	return true
}

func dfsForMaxGradeMemo(matrix [][]int, node int, currentGrade int, visited []bool, maxGrade *int, memo map[int]int) {
	n := len(matrix)
	if cachedGrade, exists := memo[node]; exists {
		currentGrade += cachedGrade
		if currentGrade > *maxGrade {
			*maxGrade = currentGrade
		}
		return
	}
	visited[node] = true
	isEnd := true

	for i := 0; i < n; i++ {
		if !visited[i] && matrix[node][i] > 0 {
			isEnd = false
			dfsForMaxGradeMemo(matrix, i, currentGrade+matrix[node][i], visited, maxGrade, memo)
		}
	}

	if isEnd {
		if currentGrade > *maxGrade {
			*maxGrade = currentGrade
		}
		memo[node] = currentGrade
	}
	visited[node] = false
}
