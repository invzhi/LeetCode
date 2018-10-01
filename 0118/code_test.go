package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	var tests = []struct {
		numRows  int
		triangle [][]int
	}{
		{
			5,
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for _, tt := range tests {
		triangle := generate(tt.numRows)
		if reflect.DeepEqual(triangle, tt.triangle) == false {
			t.Errorf("generate(%v) return %v, want %v", tt.numRows, triangle, tt.triangle)
		}
	}
}
