package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		nums     []int
		elements []int
		length   int
	}{
		{[]int{}, []int{}, 0},
		{[]int{1, 1, 2}, []int{1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		length := removeDuplicates(tt.nums)
		if length != tt.length || reflect.DeepEqual(tt.nums[:length], tt.elements) == false {
			t.Errorf("removeDuplicates(%v) return %v, %v, want %v, %v", nums, tt.nums[:length], length, tt.elements, tt.length)
		}
	}
}
