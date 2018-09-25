package leetcode

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	var tests = []struct {
		tree  *TreeNode
		order [][]int
	}{
		{
			newTree(3, 9, 20, nil, nil, 15, 7),
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			newTree(1, 2, 2, 3, nil, 3, 3, 4),
			[][]int{
				{1},
				{2, 2},
				{3, 3, 3},
				{4},
			},
		},
		{nil, nil},
	}

	for _, tt := range tests {
		order := levelOrder(tt.tree)
		if reflect.DeepEqual(order, tt.order) == false {
			t.Errorf("levelOrder(%v) return %v, want %v", tt.tree, order, tt.order)
		}
	}
}
