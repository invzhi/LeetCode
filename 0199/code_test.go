package leetcode

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		vals []int
	}{
		{newTree(1, 2, 3, nil, 5, nil, 4), []int{1, 3, 4}},
		{newTree(1, 2, 3, 4, 5, nil, 6, nil, 7, nil, nil, 8, nil, nil, 9, nil, nil, nil, 10), []int{1, 3, 6, 8, 9, 10}},
	}

	for _, tt := range tests {
		vals := rightSideView(tt.tree)
		if reflect.DeepEqual(vals, tt.vals) == false {
			t.Errorf("rightSideView(%v) return %v, want %v", tt.tree, vals, tt.vals)
		}
	}
}
