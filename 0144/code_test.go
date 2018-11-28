package leetcode

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		vals []int
	}{
		{newTree(1, nil, 2, 3), []int{1, 2, 3}},
		{newTree(1, 2, 5, 3, 4, 6, 7), []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, tt := range tests {
		vals := preorderTraversal(tt.tree)
		if reflect.DeepEqual(vals, tt.vals) == false {
			t.Errorf("preorderTraversal(%v) return %v, want %v", tt.tree, vals, tt.vals)
		}
	}
}
