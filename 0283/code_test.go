package leetcode

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var tests = []struct {
		nums  []int
		moved []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		moveZeroes(tt.nums)
		if reflect.DeepEqual(tt.nums, tt.moved) == false {
			t.Errorf("moveZeroes(%v) return %v, want %v", nums, tt.nums, tt.moved)
		}
	}
}
