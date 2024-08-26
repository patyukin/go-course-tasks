package ex1

import (
	"math"
	"reflect"
	"testing"
)

func TestBFS2(t *testing.T) {
	tests := []struct {
		name     string
		graph    [][]int
		start    int
		expected []int
	}{
		{
			name: "Simple graph",
			graph: [][]int{
				{0, 1, 4, 0, 0, 0},
				{1, 0, 4, 2, 7, 0},
				{4, 4, 0, 3, 5, 0},
				{0, 2, 3, 0, 4, 6},
				{0, 7, 5, 4, 0, 7},
				{0, 0, 0, 6, 7, 0},
			},
			start:    0,
			expected: []int{0, 1, 4, 3, 7, 9},
		},
		{
			name: "Graph with no edges",
			graph: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			start:    0,
			expected: []int{0, math.MaxInt32, math.MaxInt32},
		},
		{
			name: "Graph with a single node",
			graph: [][]int{
				{0},
			},
			start:    0,
			expected: []int{0},
		},
		{
			name: "Graph with multiple paths",
			graph: [][]int{
				{0, 3, 1, 0, 0},
				{3, 0, 1, 0, 5},
				{1, 1, 0, 7, 0},
				{0, 0, 7, 0, 2},
				{0, 5, 0, 2, 0},
			},
			start:    0,
			expected: []int{0, 2, 1, 8, 7},
		},
		{
			name: "Graph with disconnected components",
			graph: [][]int{
				{0, 2, 0, 0, 0},
				{2, 0, 3, 0, 0},
				{0, 3, 0, 0, 0},
				{0, 0, 0, 0, 1},
				{0, 0, 0, 1, 0},
			},
			start:    0,
			expected: []int{0, 2, 5, math.MaxInt32, math.MaxInt32},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BFS2(tt.graph, tt.start)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BFS() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
