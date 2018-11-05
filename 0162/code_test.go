package leetcode

import "testing"

func TestFindPeakElement(t *testing.T) {
	var tests = []struct {
		nums []int
		peak int
	}{
		{[]int{1}, 0},
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
	}

	for _, tt := range tests {
		peak := findPeakElement(tt.nums)
		if peak != tt.peak {
			t.Errorf("findPeakElement(%v) return %v, want %v", tt.nums, peak, tt.peak)
		}
	}
}
