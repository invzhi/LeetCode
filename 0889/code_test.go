package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructFromPrePost(t *testing.T) {
	var tests = []struct {
		pre  []int
		post []int
		tree *TreeNode
	}{
		{[]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}, newTree(1, 2, 3, 4, 5, 6, 7)},
		{[]int{1, 2, 3, 4}, []int{3, 4, 2, 1}, newTree(1, 2, nil, 3, 4)},
	}

	for _, tt := range tests {
		tree := constructFromPrePost(tt.pre, tt.post)
		if reflect.DeepEqual(tree, tt.tree) == false {
			t.Errorf("constructFromPrePost(%v, %v) return %v, want %v", tt.pre, tt.post, tree, tt.tree)
		}
	}
}
