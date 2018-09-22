package leetcode

import "testing"

func TestFindMin(t *testing.T) {
	var tests = []struct {
		array []int
		min   int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{2, 3, 4, 5, 1}, 1},
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 1, 2, 3}, 1},
		{[]int{5, 1, 2, 3, 4}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	}

	for _, tt := range tests {
		min := findMin(tt.array)
		if min != tt.min {
			t.Errorf("findMin(%v) return %v, want %v", tt.array, min, tt.min)
		}
	}
}
