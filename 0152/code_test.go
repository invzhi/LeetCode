package leetcode

import "testing"

func TestMaxProduct(t *testing.T) {
	var tests = []struct {
		nums []int
		max  int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{4, -2, 2, 3}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, -4, -1}, 8},
		{[]int{-2}, -2},
	}

	for _, tt := range tests {
		max := maxProduct(tt.nums)
		if max != tt.max {
			t.Errorf("maxProduct(%v) return %v, want %v", tt.nums, max, tt.max)
		}
	}
}
