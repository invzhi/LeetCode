package leetcode

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	var tests = []struct {
		A [][]int
		T [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			[][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
	}

	for _, tt := range tests {
		T := transpose(tt.A)
		if reflect.DeepEqual(T, tt.T) == false {
			t.Errorf("transpose(%v) return %v, want %v", tt.A, T, tt.T)
		}
	}
}
