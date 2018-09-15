package leetcode

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	var tests = []struct {
		grid    [][]int
		maxArea int
	}{
		{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			}, 6,
		},
		{
			[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			}, 4,
		},
		{
			[][]int{{0, 0, 0, 0, 0, 0, 0, 0}}, 0,
		},
	}

	for _, tt := range tests {
		maxArea := maxAreaOfIsland(tt.grid)
		if maxArea != tt.maxArea {
			t.Errorf("maxAreaOfIsland(%v) return %v, want %v", tt.grid, maxArea, tt.maxArea)
		}
	}
}
