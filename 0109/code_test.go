package leetcode

import (
	"reflect"
	"testing"
)

func TestSortedListToBST(t *testing.T) {
	var tests = []struct {
		list []int
		tree *TreeNode
	}{
		{[]int{-10, -3, 0, 5, 9}, newTree(0, -3, 9, -10, nil, 5)},
	}

	for _, tt := range tests {
		tree := sortedListToBST(newList(tt.list...))
		if reflect.DeepEqual(tree, tt.tree) == false {
			t.Errorf("sortedListToBST(%v) return %v, want %v", newList(tt.list...), tree, tt.tree)
		}
	}
}
