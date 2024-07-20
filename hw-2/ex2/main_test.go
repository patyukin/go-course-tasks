package main

import (
	"reflect"
	"testing"
)

func Test_calcVotes(t *testing.T) {
	type args struct {
		votes []string
	}
	tests := []struct {
		name string
		args args
		want []Candidate
	}{
		{
			name: "test1",
			args: args{
				votes: []string{"A", "B", "C"},
			},
			want: []Candidate{
				{Name: "A", Votes: 1},
				{Name: "B", Votes: 1},
				{Name: "C", Votes: 1},
			},
		},
		{
			name: "test2",
			args: args{
				votes: []string{"A", "B", "C", "A"},
			},
			want: []Candidate{
				{Name: "A", Votes: 2},
				{Name: "B", Votes: 1},
				{Name: "C", Votes: 1},
			},
		},
		{
			name: "test3",
			args: args{
				votes: []string{"A", "B", "C", "A", "B"},
			},
			want: []Candidate{
				{Name: "A", Votes: 2},
				{Name: "B", Votes: 2},
				{Name: "C", Votes: 1},
			},
		},
		{
			name: "test4",
			args: args{
				votes: []string{"A", "B", "C", "A", "B", "C"},
			},
			want: []Candidate{
				{Name: "A", Votes: 2},
				{Name: "B", Votes: 2},
				{Name: "C", Votes: 2},
			},
		},
		{
			name: "test5",
			args: args{
				votes: []string{},
			},
			want: []Candidate{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcVotes(tt.args.votes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcVotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
