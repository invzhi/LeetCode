package leetcode

import (
	"reflect"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	var tests = []struct {
		A [][]int
		B [][]int
	}{
		{
			[][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			[][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		{
			[][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 0},
			},
			[][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 1},
				{1, 0, 1, 0},
			},
		},
	}

	for _, tt := range tests {
		A := make([][]int, len(tt.A))
		for i := range A {
			A[i] = make([]int, len(tt.A[i]))
			copy(A[i], tt.A[i])
		}

		B := flipAndInvertImage(tt.A)
		if reflect.DeepEqual(B, tt.B) == false {
			t.Errorf("flipAndInvertImage(%v) return %v, want %v", A, B, tt.B)
		}
	}
}
