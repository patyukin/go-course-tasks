package ex1

import (
	"math"
)

func BFS2(graph [][]int, start int) []int {
	n := len(graph)
	distances := make([]int, n)
	for i := range distances {
		distances[i] = math.MaxInt32
	}

	distances[start] = 0

	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for i := 0; i < n; i++ {
			if graph[node][i] > 0 && distances[node]+graph[node][i] < distances[i] {
				distances[i] = distances[node] + graph[node][i]
				queue = append(queue, i)
			}
		}
	}

	return distances
}
