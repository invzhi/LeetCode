package leetcode

import (
	"reflect"
	"testing"
)

func TestIncreasingBST(t *testing.T) {
	var tests = []struct {
		vals []interface{}
		tree *TreeNode
	}{
		{[]interface{}{5, 3, 6, 2, 4, nil, 8, 1, nil, nil, nil, 7, 9}, newTree(1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6, nil, 7, nil, 8, nil, 9)},
	}

	for _, tt := range tests {
		tree := increasingBST(newTree(tt.vals...))
		if reflect.DeepEqual(tree, tt.tree) == false {
			t.Errorf("increasingBST(%v) return %v, want %v", newTree(tt.vals...), tree, tt.tree)
		}
	}
}
