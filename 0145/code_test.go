package leetcode

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		vals []int
	}{
		{newTree(1, nil, 2, 3), []int{3, 2, 1}},
		{newTree(7, 3, 6, 1, 2, 4, 5), []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, tt := range tests {
		vals := postorderTraversal(tt.tree)
		if reflect.DeepEqual(vals, tt.vals) == false {
			t.Errorf("postorderTraversal(%v) return %v, want %v", tt.tree, vals, tt.vals)
		}
	}
}
