package main

/*
BenchmarkDFSForMaxGradeIterative
BenchmarkDFSForMaxGradeIterative-12        	 9064687	       110.3 ns/op
BenchmarkDFSForMaxGradeIterativeMemo
BenchmarkDFSForMaxGradeIterativeMemo-12    	 8191459	       146.6 ns/op
BenchmarkDFSForMaxGrade
BenchmarkDFSForMaxGrade-12                 	25034204	        48.02 ns/op
BenchmarkDFSForMaxGradeMemo
BenchmarkDFSForMaxGradeMemo-12             	14624158	        80.73 ns/op
*/

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvalSequence(t *testing.T) {
	type args struct {
		mtx        [][]int
		userAnswer []int
	}

	mtx1 := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	mtx2 := [][]int{
		{0, 2, 0, 0},
		{2, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 0},
	}
	mtx3 := [][]int{
		{0, 1, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 1, 0, 0, 1},
		{0, 0, 1, 1, 0},
	}

	mtx4 := [][]int{
		{0, 3, 0, 4, 0},
		{3, 0, 5, 0, 0},
		{0, 5, 0, 0, 2},
		{4, 0, 0, 0, 1},
		{0, 0, 2, 1, 0},
	}

	mtx5 := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}

	mtx6 := [][]int{
		{0, 2, 0, 2, 0},
		{2, 0, 2, 0, 2},
		{0, 2, 0, 2, 0},
		{2, 0, 2, 0, 2},
		{0, 2, 0, 2, 0},
	}

	mtx7 := [][]int{
		{0, 1, 0, 0, 1},
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 1, 0},
	}

	mtx8 := [][]int{
		{0, 3, 0, 0, 0},
		{3, 0, 4, 0, 0},
		{0, 4, 0, 5, 0},
		{0, 0, 5, 0, 6},
		{0, 0, 0, 6, 0},
	}

	mtx9 := [][]int{
		{0, 1, 1, 1, 1},
		{1, 0, 1, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 0, 1},
		{1, 1, 1, 1, 0},
	}

	mtx10 := [][]int{
		{0, 10, 0, 0, 0},
		{10, 0, 20, 0, 0},
		{0, 20, 0, 30, 0},
		{0, 0, 30, 0, 40},
		{0, 0, 0, 40, 0},
	}

	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{
			name: "mtx 5 verticals 100%",
			args: args{
				mtx:        mtx1,
				userAnswer: []int{4, 1, 0, 2},
			},
			want: 100,
			err:  nil,
		},
		{
			name: "mtx 5 verticals 0%",
			args: args{
				mtx:        mtx1,
				userAnswer: []int{},
			},
			want: 0,
			err:  nil,
		},
		{
			name: "mtx 1 verticals 50%",
			args: args{
				mtx:        mtx1,
				userAnswer: []int{4, 1, 0},
			},
			want: 50,
			err:  nil,
		},
		{
			name: "mtx 2 vertices disconnected",
			args: args{
				mtx:        mtx2,
				userAnswer: []int{0, 1, 2},
			},
			want: 0,
			err:  errGraphIsNotConnected,
		},
		{
			name: "mtx 3 vertices fully connected",
			args: args{
				mtx:        mtx3,
				userAnswer: []int{0, 1, 2, 4, 3},
			},
			want: 100,
			err:  nil,
		},
		{
			name: "mtx 4 vertices complex",
			args: args{
				mtx:        mtx4,
				userAnswer: []int{0, 3, 4, 2, 1},
			},
			want: 85,
			err:  nil,
		},
		{
			name: "mtx 5 vertices zero weights",
			args: args{
				mtx:        mtx5,
				userAnswer: []int{2, 3},
			},
			want: 0,
			err:  errGraphIsNotConnected,
		},
		{
			name: "mtx 6 vertices with cycles",
			args: args{
				mtx:        mtx6,
				userAnswer: []int{0, 3, 2, 1, 4},
			},
			want: 100,
			err:  nil,
		},
		{
			name: "mtx 7 vertices alternate",
			args: args{
				mtx:        mtx7,
				userAnswer: []int{0, 4, 3, 2, 1},
			},
			want: 100,
			err:  nil,
		},
		{
			name: "mtx 8 vertices increasing weights",
			args: args{
				mtx:        mtx8,
				userAnswer: []int{0, 1, 2, 3, 4},
			},
			want: 100,
			err:  nil,
		},
		{
			name: "mtx 9 vertices complete graph",
			args: args{
				mtx:        mtx9,
				userAnswer: []int{0, 1, 2, 3, 4},
			},
			want: 100,
			err:  nil,
		},
		{
			name: "mtx 10 vertices varying weights",
			args: args{
				mtx:        mtx10,
				userAnswer: []int{0, 1, 2, 3, 4},
			},
			want: 100,
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalSequence(tt.args.mtx, tt.args.userAnswer)
			if !errors.Is(err, tt.err) {
				t.Errorf("got %v, want %v", err, tt.err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkDFSForMaxGradeIterative(b *testing.B) {
	matrix := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	for i := 0; i < b.N; i++ {
		maxGrade := 0
		dfsForMaxGradeIterative(matrix, 0, &maxGrade)
	}
}

func BenchmarkDFSForMaxGradeIterativeMemo(b *testing.B) {
	matrix := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	for i := 0; i < b.N; i++ {
		maxGrade := 0
		memo := make(map[int]int)
		dfsForMaxGradeIterativeMemo(matrix, 0, &maxGrade, memo)
	}
}

func BenchmarkDFSForMaxGrade(b *testing.B) {
	matrix := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	for i := 0; i < b.N; i++ {
		maxGrade := 0
		visited := make([]bool, len(matrix))
		dfsForMaxGrade(matrix, 0, 0, visited, &maxGrade)
	}
}

func BenchmarkDFSForMaxGradeMemo(b *testing.B) {
	matrix := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	for i := 0; i < b.N; i++ {
		maxGrade := 0
		visited := make([]bool, len(matrix))
		memo := make(map[int]int)
		dfsForMaxGradeMemo(matrix, 0, 0, visited, &maxGrade, memo)
	}
}

func TestDFSForMaxGrade(t *testing.T) {
	tests := []struct {
		name          string
		matrix        [][]int
		startNode     int
		expectedGrade int
	}{
		{
			name: "Test case 1 - simple graph",
			matrix: [][]int{
				{0, 2, 3, 0, 0},
				{2, 0, 0, 1, 1},
				{3, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{0, 1, 0, 0, 0},
			},
			startNode:     2,
			expectedGrade: 6, // Path: 2 -> 0 (3) -> 1 (5) -> 3 (6) = 3
		},
		{
			name: "Test case 2 - disconnected graph",
			matrix: [][]int{
				{0, 2, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 1},
				{0, 0, 1, 0},
			},
			startNode:     0,
			expectedGrade: 2, // Path: 0 -> 1 (2)
		},
		{
			name: "Test case 3 - fully connected graph",
			matrix: [][]int{
				{0, 1, 1, 1, 1},
				{1, 0, 1, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 1, 0, 1},
				{1, 1, 1, 1, 0},
			},
			startNode:     0,
			expectedGrade: 4, // Path: 0 -> 1 (1) -> 2 (1) -> 3 (1) -> 4 (1) = 4
		},
		{
			name: "Test case 4 - graph with cycles",
			matrix: [][]int{
				{0, 2, 0, 2, 0},
				{2, 0, 2, 0, 2},
				{0, 2, 0, 2, 0},
				{2, 0, 2, 0, 2},
				{0, 2, 0, 2, 0},
			},
			startNode:     0,
			expectedGrade: 8, // Path: 0 -> 3 (2) -> 2 (2) -> 1 (2) -> 4 (2) = 8
		},
		{
			name: "Test case 5 - linear graph",
			matrix: [][]int{
				{0, 3, 0, 0, 0},
				{3, 0, 4, 0, 0},
				{0, 4, 0, 5, 0},
				{0, 0, 5, 0, 6},
				{0, 0, 0, 6, 0},
			},
			startNode:     0,
			expectedGrade: 18, // Path: 0 -> 1 (3) -> 2 (4) -> 3 (5) -> 4 (6) = 18
		},
		{
			name: "Test case 6 - all zero weights",
			matrix: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			startNode:     0,
			expectedGrade: 0, // No path
		},
		{
			name: "Test case 7 - single node",
			matrix: [][]int{
				{0},
			},
			startNode:     0,
			expectedGrade: 0, // Single node, no edges
		},
		{
			name: "Test case 8 - complex graph",
			matrix: [][]int{
				{0, 10, 0, 0, 0},
				{10, 0, 20, 0, 0},
				{0, 20, 0, 30, 0},
				{0, 0, 30, 0, 40},
				{0, 0, 0, 40, 0},
			},
			startNode:     0,
			expectedGrade: 100, // Path: 0 -> 1 (10) -> 2 (20) -> 3 (30) -> 4 (40) = 100
		},
		{
			name: "Test case 9 - branching graph",
			matrix: [][]int{
				{0, 2, 2, 2, 0},
				{2, 0, 0, 0, 5},
				{2, 0, 0, 0, 4},
				{2, 0, 0, 0, 3},
				{0, 5, 4, 3, 0},
			},
			startNode:     0,
			expectedGrade: 11, // Path: 0 -> 1 (2) -> 3 (5) -> 4 (4) = 11
		},
		{
			name: "Test case 10 - another complex graph",
			matrix: [][]int{
				{0, 2, 0, 6, 0},
				{2, 0, 3, 8, 5},
				{0, 3, 0, 0, 7},
				{6, 8, 0, 0, 9},
				{0, 5, 7, 9, 0},
			},
			startNode:     0,
			expectedGrade: 26, // Path: 0 -> 3 (6) -> 1 (8) -> 4 (5) -> 2 (7) = 26
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			visited := make([]bool, len(tt.matrix))
			maxGrade := 0
			dfsForMaxGrade(tt.matrix, tt.startNode, 0, visited, &maxGrade)
			if maxGrade != tt.expectedGrade {
				t.Errorf("got %d, want %d", maxGrade, tt.expectedGrade)
			}
		})
	}
}

func TestCalcUserGrade(t *testing.T) {
	tests := []struct {
		name          string
		matrix        [][]int
		userAnswer    []int
		expectedGrade int
	}{
		{
			name: "Test case 1 - simple path",
			matrix: [][]int{
				{0, 2, 3, 0, 0},
				{2, 0, 0, 1, 1},
				{3, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{0, 1, 0, 0, 0},
			},
			userAnswer:    []int{0, 1, 3},
			expectedGrade: 3, // Path: 0 -> 1 (2) -> 3 (1) = 3
		},
		{
			name: "Test case 2 - disconnected path",
			matrix: [][]int{
				{0, 2, 0, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 1},
				{0, 0, 1, 0},
			},
			userAnswer:    []int{0, 1, 3},
			expectedGrade: 2, // Path: 0 -> 1 (2), 1 -> 3 (0) = 2
		},
		{
			name: "Test case 3 - fully connected graph",
			matrix: [][]int{
				{0, 1, 1, 1, 1},
				{1, 0, 1, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 1, 0, 1},
				{1, 1, 1, 1, 0},
			},
			userAnswer:    []int{0, 1, 2, 3, 4},
			expectedGrade: 4, // Path: 0 -> 1 (1) -> 2 (1) -> 3 (1) -> 4 (1) = 4
		},
		{
			name: "Test case 4 - graph with cycles",
			matrix: [][]int{
				{0, 2, 0, 2, 0},
				{2, 0, 2, 0, 2},
				{0, 2, 0, 2, 0},
				{2, 0, 2, 0, 2},
				{0, 2, 0, 2, 0},
			},
			userAnswer:    []int{0, 3, 4, 1},
			expectedGrade: 6, // Path: 0 -> 3 (2) -> 4 (2) -> 1 (2) = 6
		},
		{
			name: "Test case 5 - linear path",
			matrix: [][]int{
				{0, 3, 0, 0, 0},
				{3, 0, 4, 0, 0},
				{0, 4, 0, 5, 0},
				{0, 0, 5, 0, 6},
				{0, 0, 0, 6, 0},
			},
			userAnswer:    []int{0, 1, 2, 3, 4},
			expectedGrade: 18, // Path: 0 -> 1 (3) -> 2 (4) -> 3 (5) -> 4 (6) = 18
		},
		{
			name: "Test case 6 - path with zero weights",
			matrix: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
			},
			userAnswer:    []int{1, 2, 3},
			expectedGrade: 2, // Path: 1 -> 2 (1) -> 3 (1) = 2
		},
		{
			name: "Test case 7 - empty path",
			matrix: [][]int{
				{0, 1, 1, 1, 1},
				{1, 0, 1, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 1, 0, 1},
				{1, 1, 1, 1, 0},
			},
			userAnswer:    []int{},
			expectedGrade: 0, // No path
		},
		{
			name: "Test case 8 - single node path",
			matrix: [][]int{
				{0, 1, 1, 1, 1},
				{1, 0, 1, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 1, 0, 1},
				{1, 1, 1, 1, 0},
			},
			userAnswer:    []int{0},
			expectedGrade: 0, // Single node, no edges
		},
		{
			name: "Test case 9 - complex path",
			matrix: [][]int{
				{0, 10, 0, 0, 0},
				{10, 0, 20, 0, 0},
				{0, 20, 0, 30, 0},
				{0, 0, 30, 0, 40},
				{0, 0, 0, 40, 0},
			},
			userAnswer:    []int{0, 1, 2, 3, 4},
			expectedGrade: 100, // Path: 0 -> 1 (10) -> 2 (20) -> 3 (30) -> 4 (40) = 100
		},
		{
			name: "Test case 10 - invalid path",
			matrix: [][]int{
				{0, 2, 3, 0, 0},
				{2, 0, 0, 1, 1},
				{3, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{0, 1, 0, 0, 0},
			},
			userAnswer:    []int{0, 4, 1},
			expectedGrade: 1, // Path: 0 -> 4 (0) -> 1 (1) = 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grade := calcUserGrade(tt.matrix, tt.userAnswer)
			if grade != tt.expectedGrade {
				t.Errorf("got %d, want %d", grade, tt.expectedGrade)
			}
		})
	}
}

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name        string
		matrix      [][]int
		userAnswer  []int
		expectValid error
	}{
		{
			name: "Valid input - fully connected graph",
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 0, 1},
				{1, 1, 1, 0},
			},
			userAnswer:  []int{0, 1, 2, 3},
			expectValid: nil,
		},
		{
			name: "Invalid input - non-square matrix",
			matrix: [][]int{
				{0, 1, 1},
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 1},
			},
			userAnswer:  []int{0, 1, 2},
			expectValid: errMatrixIsNotSquare,
		},
		{
			name: "Invalid input - loops in matrix",
			matrix: [][]int{
				{1, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 0, 1},
				{1, 1, 1, 0},
			},
			userAnswer:  []int{0, 1, 2, 3},
			expectValid: errMatrixContainsLoops,
		},
		{
			name: "Invalid input - asymmetric matrix",
			matrix: [][]int{
				{0, 1, 0, 1},
				{1, 0, 1, 1},
				{0, 2, 0, 1},
				{1, 1, 1, 0},
			},
			userAnswer:  []int{0, 1, 2, 3},
			expectValid: errMatrixIsNotSymmetrical,
		},
		{
			name: "Invalid input - negative weights",
			matrix: [][]int{
				{0, 1, -1, 1},
				{1, 0, 1, 1},
				{-1, 1, 0, 1},
				{1, 1, 1, 0},
			},
			userAnswer:  []int{0, 1, 2, 3},
			expectValid: errMatrixContainsNegativeValues,
		},
		{
			name: "Invalid input - all zero weights",
			matrix: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			userAnswer:  []int{0, 1, 2, 3},
			expectValid: errOnlyZeroValues,
		},
		{
			name: "Invalid input - duplicate answers",
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 0, 1},
				{1, 1, 1, 0},
			},
			userAnswer:  []int{0, 1, 1, 3},
			expectValid: errRepeatedAnswers,
		},
		{
			name: "Invalid input - answers out of range",
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 0, 1},
				{1, 1, 1, 0},
			},
			userAnswer:  []int{0, 1, 4, 3},
			expectValid: errPutOfRangeValues,
		},
		{
			name: "Invalid input - disconnected graph",
			matrix: [][]int{
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 1},
				{0, 0, 1, 0},
			},
			userAnswer:  []int{0, 1, 2, 3},
			expectValid: errGraphIsNotConnected,
		},
		{
			name: "Invalid input - bad graph",
			matrix: [][]int{
				{0, 1, 3, 0},
				{1, 0, 2, 0},
				{0, 3, 0, 1},
				{0, 0, 1, 0},
			},
			userAnswer:  []int{0, 1, 2, 3},
			expectValid: errMatrixIsNotSymmetrical,
		},
		{
			name: "Invalid input - not symmetric graph",
			matrix: [][]int{
				{0, 1, 3, 0},
				{1, 0, 2, 0},
				{0, 3, 0, 1},
				{0, 9, 1, 0},
			},
			userAnswer:  []int{0, 1, 2, 3},
			expectValid: errMatrixIsNotSymmetrical,
		},
		{
			name: "Invalid input - not connected graph",
			matrix: [][]int{
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			userAnswer:  []int{0, 1, 2, 3},
			expectValid: errGraphIsNotConnected,
		},
		{
			name:        "Valid input - single node",
			matrix:      [][]int{{0}},
			userAnswer:  []int{0},
			expectValid: errOnlyZeroValues,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := validateInput(tt.matrix, tt.userAnswer)
			if !errors.Is(valid, tt.expectValid) {
				t.Errorf("got %v, want %v", valid, tt.expectValid)
			}
		})
	}
}
