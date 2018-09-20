package leetcode

import "testing"

func TestMissingNumber(t *testing.T) {
	var tests = []struct {
		nums    []int
		missing int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, tt := range tests {
		missing := missingNumber(tt.nums)
		if missing != tt.missing {
			t.Errorf("missingNumber(%v) return %v, want %v", tt.nums, missing, tt.missing)
		}
	}
}
