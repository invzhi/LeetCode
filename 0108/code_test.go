package leetcode

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	var tests = []struct {
		array []int
		tree  *TreeNode
	}{
		{[]int{-10, -3, 0, 5, 9}, newTree(0, -3, 9, -10, nil, 5)},
	}

	for _, tt := range tests {
		tree := sortedArrayToBST(tt.array)
		if reflect.DeepEqual(tree, tt.tree) == false {
			t.Errorf("sortedArrayToBST(%v) return %v, want %v", tt.array, tree, tt.tree)
		}
	}
}
