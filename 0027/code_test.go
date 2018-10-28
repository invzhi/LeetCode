package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		nums  []int
		val   int
		after []int
	}{
		{[]int{3, 2, 2, 3}, 3, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		length := removeElement(tt.nums, tt.val)
		if reflect.DeepEqual(tt.nums[:length], tt.after) == false {
			t.Errorf("removeElement(%v, %v) return %v, want %v", nums, tt.val, tt.nums[:length], tt.after)
		}
	}
}
