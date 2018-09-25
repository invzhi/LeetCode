package leetcode

import "testing"

func TestIsBalanced(t *testing.T) {
	var tests = []struct {
		tree     *TreeNode
		balanced bool
	}{
		{newTree(3, 9, 20, nil, nil, 15, 7), true},
		{newTree(1, 2, 2, 3), true},
		{newTree(1, 2, 2, 3, 3, nil, nil, 4, 4), false},
		{newTree(1, 2, 2, 3, nil, 3, 3, 4), false},
		{newTree(1, 2, 2, nil, nil, nil, 3, 4, 4), false},
	}

	for _, tt := range tests {
		balanced := isBalanced(tt.tree)
		if balanced != tt.balanced {
			t.Errorf("isBalanced(%v) return %v, want %v", tt.tree, balanced, tt.balanced)
		}
	}
}
