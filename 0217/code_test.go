package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {
	var tests = []struct {
		nums      []int
		duplicate bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tt := range tests {
		duplicate := containsDuplicate(tt.nums)
		if duplicate != tt.duplicate {
			t.Errorf("containsDuplicate(%v) return %v, want %v", tt.nums, duplicate, tt.duplicate)
		}
	}
}
