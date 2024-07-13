package ex1

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		adjMatrix [][]int
		start     int
		want      []int
	}{
		{
			adjMatrix: [][]int{
				{0, 1, 1, 0, 0},
				{1, 0, 1, 1, 0},
				{1, 1, 0, 0, 1},
				{0, 1, 0, 0, 1},
				{0, 0, 1, 1, 0},
			},
			start: 0,
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			adjMatrix: [][]int{
				{0, 1, 0, 0, 0},
				{1, 0, 1, 1, 0},
				{0, 1, 0, 0, 1},
				{0, 1, 0, 0, 1},
				{0, 0, 1, 1, 0},
			},
			start: 1,
			want:  []int{1, 0, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run("BFS Test", func(t *testing.T) {
			got := BFS(tt.adjMatrix, tt.start)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
