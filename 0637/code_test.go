package leetcode

import (
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	var tests = []struct {
		tree    *TreeNode
		average []float64
	}{
		{newTree(3, 9, 20, nil, nil, 15, 7), []float64{3, 14.5, 11}},
	}

	for _, tt := range tests {
		average := averageOfLevels(tt.tree)
		if reflect.DeepEqual(average, tt.average) == false {
			t.Errorf("averageOfLevels(%v) return %v, want %v", tt.tree, average, tt.average)
		}
	}
}
