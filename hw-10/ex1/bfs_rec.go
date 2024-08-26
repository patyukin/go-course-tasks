package ex1

func BFS3(graph [][]int, dist []int, queue []int) {
	if len(queue) == 0 {
		return
	}

	node := queue[0]
	queue = queue[1:]

	for i := 0; i < len(graph[node]); i++ {
		if graph[node][i] > 0 && dist[node]+graph[node][i] < dist[i] {
			dist[i] = dist[node] + graph[node][i]
			queue = append(queue, i)
		}
	}

	BFS3(graph, dist, queue)
}
