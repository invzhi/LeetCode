package leetcode

import "testing"

func TestIslandPerimeter(t *testing.T) {
	var tests = []struct {
		grid      [][]int
		perimeter int
	}{
		{
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			}, 16,
		},
	}

	for _, tt := range tests {
		perimeter := islandPerimeter(tt.grid)
		if perimeter != tt.perimeter {
			t.Errorf("islandPerimeter(%v) return %v, want %v", tt.grid, perimeter, tt.perimeter)
		}
	}
}
