package leetcode

import (
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		sum  int
		path [][]int
	}{
		{newTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1), 22, [][]int{[]int{5, 4, 11, 2}, []int{5, 8, 4, 5}}},
	}

	for _, tt := range tests {
		path := pathSum(tt.tree, tt.sum)
		if reflect.DeepEqual(path, tt.path) == false {
			t.Errorf("PathSum(%v, %v) return %v, want %v", tt.tree, tt.sum, path, tt.path)
		}
	}
}
