package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	var tests = []struct {
		nums []int
		max  int
	}{
		{[]int{-2, -1}, -1},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, tt := range tests {
		max := maxSubArray(tt.nums)
		if max != tt.max {
			t.Errorf("maxSubArray(%v) return %v, want %v", tt.nums, max, tt.max)
		}
	}
}
