package leetcode

import "testing"

func TestMaxDepth(t *testing.T) {
	var tests = []struct {
		tree  *TreeNode
		depth int
	}{
		{newTree(3, 9, 20, nil, nil, 15, 7), 3},
		{newTree(1, 2, 2, 3), 3},
		{newTree(1, 2, 2, 3, 3, nil, nil, 4, 4), 4},
		{newTree(1, 2, 2, 3, nil, 3, 3, 4), 4},
		{newTree(1, 2, 2, nil, nil, nil, 3, 4, 4), 4},
	}

	for _, tt := range tests {
		depth := maxDepth(tt.tree)
		if depth != tt.depth {
			t.Errorf("maxDepth(%v) return %v, want %v", tt.tree, depth, tt.depth)
		}
	}
}
