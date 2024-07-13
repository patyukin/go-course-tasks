package ex1

func BFS(adjMatrix [][]int, start int) []int {
	n := len(adjMatrix)
	visited := make([]bool, n)
	order := make([]int, 0)

	queue := make([]int, 0)
	queue = append(queue, start)
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)

		for neighbor := 0; neighbor < n; neighbor++ {
			if adjMatrix[node][neighbor] == 1 && !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return order
}
