package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		index  int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{nil, 0, 0},
	}

	for _, tt := range tests {
		index := searchInsert(tt.nums, tt.target)
		if index != tt.index {
			t.Errorf("searchInsert(%v, %v) return %v, want %v", tt.nums, tt.target, index, tt.index)
		}
	}
}
