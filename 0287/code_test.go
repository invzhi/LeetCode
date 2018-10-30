package leetcode

import "testing"

func TestFindDuplicate(t *testing.T) {
	var tests = []struct {
		nums []int
		dup  int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
	}

	for _, tt := range tests {
		dup := findDuplicate(tt.nums)
		if dup != tt.dup {
			t.Errorf("findDuplicate(%v) return %v, want %v", tt.nums, dup, tt.dup)
		}
	}
}
