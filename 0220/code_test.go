package leetcode

import "testing"

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	var tests = []struct {
		nums      []int
		k         int
		t         int
		duplicate bool
	}{
		{[]int{-1, -1}, 1, -1, false},
		{[]int{-3, 3}, 2, 4, false},
		{[]int{2, 1}, 1, 1, true},
		{[]int{1, 2, 3, 1}, 3, 0, true},
		{[]int{1, 0, 1, 1}, 1, 2, true},
		{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
	}

	for _, tt := range tests {
		duplicate := containsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t)
		if duplicate != tt.duplicate {
			t.Errorf("containsNearbyAlmostDuplicate(%v, %v, %v) return %v, want %v", tt.nums, tt.k, tt.t, duplicate, tt.duplicate)
		}
	}
}
