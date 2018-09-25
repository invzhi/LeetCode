package leetcode

import "testing"

func TestLongestUnivaluePath(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		length int
	}{
		{newTree(5, 4, 5, 1, 1, nil, 5), 2},
		{newTree(5, 5, 4, nil, 5, 1, 1), 2},
		{newTree(1, 4, 5, 4, 4, nil, 5), 2},
		{newTree(1, 5, 4, 5, nil, 4, 4), 2},
		{newTree(26, 26, 26), 2},
		{newTree(26, 26, 26, 26, 26), 3},
	}

	for _, tt := range tests {
		length := longestUnivaluePath(tt.tree)
		if length != tt.length {
			t.Errorf("longestUnivaluePath(%v) return %v, want %v", tt.tree, length, tt.length)
		}
	}
}
