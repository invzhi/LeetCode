package leetcode

import "testing"

func TestNumIslands(t *testing.T) {
	var tests = []struct {
		grid [][]byte
		num  int
	}{
		{
			[][]byte{
				[]byte("11110"),
				[]byte("11010"),
				[]byte("11000"),
				[]byte("00000"),
			}, 1,
		},
		{
			[][]byte{
				[]byte("11000"),
				[]byte("11000"),
				[]byte("00100"),
				[]byte("00011"),
			}, 3,
		},
	}

	for _, tt := range tests {
		num := numIslands(tt.grid)
		if num != tt.num {
			t.Errorf("numIslands(%s) return %v, want %v", tt.grid, num, tt.num)
		}
	}
}
