package leetcode

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	var tests = []struct {
		nums         []int
		permutations [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{2, 2, 1, 1}, [][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		permutations := permuteUnique(tt.nums)
		if reflect.DeepEqual(permutations, tt.permutations) == false {
			t.Errorf("permuteUnique(%v) return %v, want %v", nums, permutations, tt.permutations)
		}
	}
}
