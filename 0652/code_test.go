package leetcode

import (
	"reflect"
	"testing"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	var tests = []struct {
		tree     *TreeNode
		subtrees []*TreeNode
	}{
		{newTree(1, 2, 3, 4, nil, 2, 4, nil, nil, 4), []*TreeNode{newTree(4), newTree(2, 4)}},
		{newTree(0, 0, 0, 0, nil, nil, 0, 0, 0, 0, 0), []*TreeNode{newTree(0), newTree(0, 0, 0)}},
	}

	for _, tt := range tests {
		subtrees := findDuplicateSubtrees(tt.tree)
		if reflect.DeepEqual(subtrees, tt.subtrees) == false {
			t.Errorf("findDuplicateSubtrees(%v) return %v, want %v", tt.tree, subtrees, tt.subtrees)
		}
	}
}
