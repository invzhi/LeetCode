package leetcode

import "testing"

func TestThreeSumClosest(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		sum    int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-1, 0, 1, 2, -1, -4, -4}, 1, 1},
	}

	for _, tt := range tests {
		sum := threeSumClosest(tt.nums, tt.target)
		if sum != tt.sum {
			t.Errorf("threeSumClosest(%v, %v) return %v, want %v", tt.nums, tt.target, sum, tt.sum)
		}
	}
}
