package leetcode

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	var tests = []struct {
		nums      []int
		k         int
		duplicate bool
	}{
		{[]int{99, 99}, 2, true},
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, tt := range tests {
		duplicate := containsNearbyDuplicate(tt.nums, tt.k)
		if duplicate != tt.duplicate {
			t.Errorf("containsNearbyDuplicate(%v, %v) return %v, want %v", tt.nums, tt.k, duplicate, tt.duplicate)
		}
	}
}
