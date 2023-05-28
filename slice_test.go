package goext

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args[T any] struct {
		sourceSlice []T
		fn          func(index int, item T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "int",
			args: args[int]{
				sourceSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(index int, item int) bool {
					return item >= 5
				},
			},
			want: []int{5, 6, 7, 8, 9},
		},
		{
			name: "int",
			args: args[int]{
				sourceSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(index int, item int) bool {
					return item >= 3
				},
			},
			want: []int{3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "int",
			args: args[int]{
				sourceSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(index int, item int) bool {
					return item <= 3
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.sourceSlice, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
	tests2 := []testCase[float64]{
		{
			name: "float",
			args: args[float64]{
				sourceSlice: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0},
				fn: func(index int, item float64) bool {
					return item >= 5.0
				},
			},
			want: []float64{5.0, 6.0, 7.0, 8.0, 9.0},
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.sourceSlice, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, R any] struct {
		sourceSlice []T
		fn          func(index int, item T) R
	}
	type testCase[T any, R any] struct {
		name string
		args args[T, R]
		want []R
	}
	tests := []testCase[int, int]{
		{
			name: "int",
			args: args[int, int]{
				sourceSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(index int, item int) int {
					return item * 1
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.sourceSlice, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args[T any, R any] struct {
		sourceSlice []T
		fn          func(index int, item T, agg R) R
		initial     R
	}
	type testCase[T any, R any] struct {
		name string
		args args[T, R]
		want R
	}
	tests := []testCase[int, int]{
		{
			name: "int",
			args: args[int, int]{
				sourceSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(index int, item int, initial int) int {
					return initial + item
				},
				initial: 0,
			},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.sourceSlice, tt.args.fn, tt.args.initial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
