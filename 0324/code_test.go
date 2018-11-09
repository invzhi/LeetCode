package leetcode

import (
	"reflect"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	var tests = []struct {
		nums   []int
		wiggle []int
	}{
		{[]int{1, 5, 1, 1, 6, 4}, []int{1, 5, 1, 4, 1, 6}},
		{[]int{1, 3, 2, 2, 3, 1}, []int{2, 3, 1, 3, 1, 2}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{3, 5, 1, 4, 2, 6}},
		{[]int{1}, []int{1}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		wiggleSort(tt.nums)
		if reflect.DeepEqual(tt.nums, tt.wiggle) == false {
			t.Errorf("wiggleSort(%v) return %v, want %v", nums, tt.nums, tt.wiggle)
		}
	}
}
