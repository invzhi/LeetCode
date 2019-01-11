package leetcode

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	var tests = []struct {
		nums     []int
		target   int
		position []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
		{[]int{}, 0, []int{-1, -1}},
	}

	for _, tt := range tests {
		position := searchRange(tt.nums, tt.target)
		if reflect.DeepEqual(position, tt.position) == false {
			t.Errorf("searchRange(%v, %v) return %v, want %v", tt.nums, tt.target, position, tt.position)
		}
	}
}
