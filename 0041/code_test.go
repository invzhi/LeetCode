package leetcode

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	var tests = []struct {
		nums    []int
		missing int
	}{
		{[]int{1, 1}, 2},
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		missing := firstMissingPositive(tt.nums)
		if missing != tt.missing {
			t.Errorf("firstMissingPositive(%v) return %v, want %v", nums, missing, tt.missing)
		}
	}
}
