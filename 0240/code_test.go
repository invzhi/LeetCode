package leetcode

import "testing"

func TestSearchMatrix(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		target int
		found  bool
	}{
		{
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			}, 5, true,
		},
		{
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			}, 20, false,
		},
		{nil, 0, false},
	}

	for _, tt := range tests {
		found := searchMatrix(tt.matrix, tt.target)
		if found != tt.found {
			t.Errorf("searchMatrix(%v, %v) return %v, want %v", tt.matrix, tt.target, found, tt.found)
		}
	}
}
