package leetcode

import "testing"

func TestMajorityElement(t *testing.T) {
	var tests = []struct {
		nums     []int
		majority int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, tt := range tests {
		majority := majorityElement(tt.nums)
		if majority != tt.majority {
			t.Errorf("majorityElement(%v) return %v, want %v", tt.nums, majority, tt.majority)
		}
	}
}
