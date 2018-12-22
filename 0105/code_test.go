package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var tests = []struct {
		preorder []int
		inorder  []int
		tree     *TreeNode
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, newTree(3, 9, 20, nil, nil, 15, 7)},
	}

	for _, tt := range tests {
		tree := buildTree(tt.preorder, tt.inorder)
		if reflect.DeepEqual(tree, tt.tree) == false {
			t.Errorf("buildTree(%v, %v) return %v, want %v", tt.preorder, tt.inorder, tree, tt.tree)
		}
	}
}
