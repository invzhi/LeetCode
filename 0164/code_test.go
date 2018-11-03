package leetcode

import "testing"

func TestMaximumGap(t *testing.T) {
	var tests = []struct {
		nums []int
		diff int
	}{
		{[]int{10}, 0},
		{[]int{1, 3, 100}, 97},
		{[]int{3, 1, 100}, 97},
		{[]int{3, 3, 2, 1}, 1},
		{[]int{3, 6, 9, 1}, 3},
	}

	for _, tt := range tests {
		diff := maximumGap(tt.nums)
		if diff != tt.diff {
			t.Errorf("maximumGap(%v) return %v, want %v", tt.nums, diff, tt.diff)
		}
	}
}
