package leetcode

import (
	"reflect"
	"testing"
)

func TestSubtreeWithAllDeepest(t *testing.T) {
	var tests = []struct {
		tree    *TreeNode
		subtree *TreeNode
	}{
		{newTree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4), newTree(2, 7, 4)},
	}

	for _, tt := range tests {
		subtree := subtreeWithAllDeepest(tt.tree)
		if reflect.DeepEqual(subtree, tt.subtree) == false {
			t.Errorf("subtreeWithAllDeepest(%v) return %v, want %v", tt.tree, subtree, tt.subtree)
		}
	}
}
