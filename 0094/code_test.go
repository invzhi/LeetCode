package leetcode

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		vals []int
	}{
		{newTree(1, nil, 2, 3), []int{1, 3, 2}},
		{newTree(4, 2, 6, 1, 3, 5, 7), []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, tt := range tests {
		vals := inorderTraversal(tt.tree)
		if reflect.DeepEqual(vals, tt.vals) == false {
			t.Errorf("inorderTraversal(%v) return %v, want %v", tt.tree, vals, tt.vals)
		}
	}
}
