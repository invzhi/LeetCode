package leetcode

import (
	"reflect"
	"testing"
)

func TestLevelOrderBttom(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		values [][]int
	}{
		{newTree(3, 9, 20, nil, nil, 15, 7), [][]int{{15, 7}, {9, 20}, {3}}},
	}

	for _, tt := range tests {
		values := levelOrderBottom(tt.tree)
		if reflect.DeepEqual(values, tt.values) == false {
			t.Errorf("levelOrderBottom(%v) return %v, want %v", tt.tree, values, tt.values)
		}
	}
}
