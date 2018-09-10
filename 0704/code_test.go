package leetcode

import "testing"

func TestSearch(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		index  int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{nil, 0, -1},
	}

	for _, tt := range tests {
		index := search(tt.nums, tt.target)
		if index != tt.index {
			t.Errorf("search(%v, %v) return %v, want %v", tt.nums, tt.target, index, tt.index)
		}
	}
}
