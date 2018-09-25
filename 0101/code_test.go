package leetcode

import "testing"

func TestIsSymmetric(t *testing.T) {
	var tests = []struct {
		tree      *TreeNode
		symmetric bool
	}{
		{newTree(1, 2, 2, 3, 4, 4, 3), true},
		{newTree(1, 2, 2, nil, 3, nil, 3), false},
		{newTree(1, 2, 2, nil, 3, 4), false},
		{nil, true},
	}

	for _, tt := range tests {
		symmetric := isSymmetric(tt.tree)
		if symmetric != tt.symmetric {
			t.Errorf("isSymmetric(%v) return %v, want %v", tt.tree, symmetric, tt.symmetric)
		}
	}
}
