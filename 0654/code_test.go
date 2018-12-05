package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructMaxiumBinaryTree(t *testing.T) {
	var tests = []struct {
		nums []int
		tree *TreeNode
	}{
		{[]int{3, 2, 1, 6, 0, 5}, newTree(6, 3, 5, nil, 2, 0, nil, nil, 1)},
	}

	for _, tt := range tests {
		tree := constructMaximumBinaryTree(tt.nums)
		if reflect.DeepEqual(tree, tt.tree) == false {
			t.Errorf("constructMaximumBinaryTree(%v) return %v, want %v", tt.nums, tree, tt.tree)
		}
	}
}
