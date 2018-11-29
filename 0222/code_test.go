package leetcode

import "testing"

func TestCountNodes(t *testing.T) {
	var tests = []struct {
		tree  *TreeNode
		count int
	}{
		{newTree(1), 1},
		{newTree(1, 2, 3, 4), 4},
		{newTree(1, 2, 3, 4, 5, 6), 6},
	}

	for _, tt := range tests {
		count := countNodes(tt.tree)
		if count != tt.count {
			t.Errorf("countNodes(%v) return %v, want %v", tt.tree, count, tt.count)
		}
	}
}
