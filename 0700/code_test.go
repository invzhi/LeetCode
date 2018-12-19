package leetcode

import (
	"reflect"
	"testing"
)

func TestSearchBST(t *testing.T) {
	var tests = []struct {
		tree    *TreeNode
		val     int
		subtree *TreeNode
	}{
		{newTree(4, 2, 7, 1, 3), 2, newTree(2, 1, 3)},
		{newTree(4, 2, 7, 1, 3), 5, nil},
	}

	for _, tt := range tests {
		subtree := searchBST(tt.tree, tt.val)
		if reflect.DeepEqual(subtree, tt.subtree) == false {
			t.Errorf("searchBST(%v, %v) return %v, want %v", tt.tree, tt.val, subtree, tt.subtree)
		}
	}
}
