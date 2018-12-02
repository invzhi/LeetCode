package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	var tests = []struct {
		n     int
		trees []*TreeNode
	}{
		{0, nil},
		{
			3,
			[]*TreeNode{
				newTree(1, nil, 2, nil, 3),
				newTree(1, nil, 3, 2),
				newTree(2, 1, 3),
				newTree(3, 1, nil, nil, 2),
				newTree(3, 2, nil, 1),
			},
		},
	}

	for _, tt := range tests {
		trees := generateTrees(tt.n)
		if reflect.DeepEqual(trees, tt.trees) == false {
			t.Errorf("generateTrees(%v) return %v, want %v", tt.n, trees, tt.trees)
		}
	}
}
