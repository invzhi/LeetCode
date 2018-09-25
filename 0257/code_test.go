package leetcode

import (
	"reflect"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	var tests = []struct {
		tree  *TreeNode
		paths []string
	}{
		{newTree(1, 2, 3, nil, 5), []string{"1->2->5", "1->3"}},
	}

	for _, tt := range tests {
		paths := binaryTreePaths(tt.tree)
		if reflect.DeepEqual(paths, tt.paths) == false {
			t.Errorf("binaryTreePaths(%v) return %v, want %v", tt.tree, paths, tt.paths)
		}
	}
}
