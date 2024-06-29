package main

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		slices [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				slices: [][]int{{1, 2, 3, 2}, {3, 2}},
			},
			want: []int{2, 3},
		},
		{
			name: "test2",
			args: args{
				slices: [][]int{{1, 2, 3, 2}, {3, 2}, {}},
			},
			want: []int{},
		},
		{
			name: "test3",
			args: args{
				slices: [][]int{{1, 2, 3, 2}},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test4",
			args: args{
				slices: [][]int{{}},
			},
			want: []int{},
		},
		{
			name: "test5",
			args: args{
				slices: [][]int{{}, {1, 2, 3, 2}},
			},
			want: []int{},
		},
		{
			name: "test6",
			args: args{
				slices: [][]int{{1, 2, 3, 2}, {3, 2}, {1, 2, 3, 2}},
			},
			want: []int{2, 3},
		},
		{
			name: "test7",
			args: args{
				slices: [][]int{{1, 2, 3, 2}, {3, 2}, {1, 2, 3, 2}, {3, 2}},
			},
			want: []int{2, 3},
		},
		{
			name: "test8",
			args: args{
				slices: [][]int{{1, 2, 3, 2}, {3, 2}, {1, 2, 3, 2}, {3, 2}, {1, 2, 3, 2}},
			},
			want: []int{2, 3},
		},
		{
			name: "test9",
			args: args{
				slices: [][]int{{1, 2, 3, 2}, {3, 2}, {1, 2, 3, 2}, {3, 2}, {1, 2, 3, 2}, {3, 2}},
			},
			want: []int{2, 3},
		},
		{
			name: "test9",
			args: args{
				slices: [][]int{{1, 2}, {2, 3}, {2, 4, 1}},
			},
			want: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.slices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
