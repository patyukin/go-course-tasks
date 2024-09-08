package ex1

import (
	"reflect"
	"testing"
)

func TestBFS3(t *testing.T) {
	tests := []struct {
		name     string
		graph    [][]int
		start    int
		expected []int
	}{
		{
			name: "Simple graph with equal weights",
			graph: [][]int{
				{0, 1, 1, 0},
				{1, 0, 1, 1},
				{1, 1, 0, 1},
				{0, 1, 1, 0},
			},
			start:    0,
			expected: []int{0, 1, 1, 2},
		},
		{
			name: "Larger graph with varied weights",
			graph: [][]int{
				{0, 4, 2, 0, 0},
				{4, 0, 5, 10, 0},
				{2, 5, 0, 3, 4},
				{0, 10, 3, 0, 1},
				{0, 0, 4, 1, 0},
			},
			start:    0,
			expected: []int{0, 4, 2, 5, 6},
		},
		{
			name: "Linear graph",
			graph: [][]int{
				{0, 1, 0, 0, 0},
				{1, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 1, 0, 1},
				{0, 0, 0, 1, 0},
			},
			start:    0,
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			name: "Triangle graph",
			graph: [][]int{
				{0, 2, 2},
				{2, 0, 1},
				{2, 1, 0},
			},
			start:    0,
			expected: []int{0, 2, 2},
		},
		{
			name: "Graph with one isolated node",
			graph: [][]int{
				{0, 3, 0, 0, 0},
				{3, 0, 5, 0, 0},
				{0, 5, 0, 2, 0},
				{0, 0, 2, 0, 7},
				{0, 0, 0, 7, 0},
			},
			start:    0,
			expected: []int{0, 3, 8, 10, 17},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distances := make([]int, len(tt.graph))
			for i := range distances {
				distances[i] = int(^uint(0) >> 1) // math.MaxInt32
			}
			distances[tt.start] = 0

			queue := []int{tt.start}
			BFS3(tt.graph, distances, queue)

			if !reflect.DeepEqual(distances, tt.expected) {
				t.Errorf("BFS() = %v, expected %v", distances, tt.expected)
			}
		})
	}
}
