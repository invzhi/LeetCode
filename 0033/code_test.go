package leetcode

import "testing"

func TestSearch(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		index  int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 2, 2},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 3, -1},
		{[]int{2, 4, 5, 6, 7, 0, 1}, 3, -1},
		{[]int{8, 9, 2, 3, 4}, 9, 1},
		{[]int{1}, 0, -1},
		{nil, 0, -1},
	}

	for _, tt := range tests {
		index := search(tt.nums, tt.target)
		if index != tt.index {
			t.Errorf("search(%v, %v) return %v, want %v", tt.nums, tt.target, index, tt.index)
		}
	}
}
