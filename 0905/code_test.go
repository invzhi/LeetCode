package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	var tests = []struct {
		A []int
		B []int
	}{
		{[]int{3, 1, 2, 4}, []int{4, 2, 1, 3}},
	}

	for _, tt := range tests {
		A := make([]int, len(tt.A))
		copy(A, tt.A)

		B := sortArrayByParity(tt.A)
		if reflect.DeepEqual(B, tt.B) == false {
			t.Errorf("sortArrayByParity(%v) return %v, want %v", A, B, tt.B)
		}
	}
}
