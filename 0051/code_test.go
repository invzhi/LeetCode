package leetcode

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	var tests = []struct {
		n         int
		solutions [][]string
	}{
		{
			4,
			[][]string{
				{
					".Q..",
					"...Q",
					"Q...",
					"..Q.",
				},
				{
					"..Q.",
					"Q...",
					"...Q",
					".Q..",
				},
			},
		},
	}

	for _, tt := range tests {
		solutions := solveNQueens(tt.n)
		if reflect.DeepEqual(solutions, tt.solutions) == false {
			t.Errorf("solveNQueens(%v) return %v, want %v", tt.n, solutions, tt.solutions)
		}
	}
}
