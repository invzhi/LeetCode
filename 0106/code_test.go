package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var tests = []struct {
		inorder   []int
		postorder []int
		tree      *TreeNode
	}{
		{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}, newTree(3, 9, 20, nil, nil, 15, 7)},
	}

	for _, tt := range tests {
		tree := buildTree(tt.inorder, tt.postorder)
		if reflect.DeepEqual(tree, tt.tree) == false {
			t.Errorf("buildTree(%v, %v) return %v, want %v", tt.inorder, tt.postorder, tree, tt.tree)
		}
	}
}
