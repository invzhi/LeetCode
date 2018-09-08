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
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			}, 3, true,
		},
		{
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			}, 13, false,
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
