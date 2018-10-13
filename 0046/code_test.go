package leetcode

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	var tests = []struct {
		nums         []int
		permutations [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {1, 3, 2}, {2, 3, 1}, {3, 1, 2}}},
	}

	for _, tt := range tests {
		permutations := permute(tt.nums)
		if reflect.DeepEqual(permutations, tt.permutations) == false {
			t.Errorf("permute(%v) return %v, want %v", tt.nums, permutations, tt.permutations)
		}
	}
}
