package leetcode

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	var tests = []struct {
		tree     *TreeNode
		diameter int
	}{
		{newTree(1, 2, 3, 4, 5), 3},
		{newTree(1, 2, 3, nil, nil, 4, 5), 3},
	}

	for _, tt := range tests {
		diameter := diameterOfBinaryTree(tt.tree)
		if diameter != tt.diameter {
			t.Errorf("diameterOfBinaryTree(%v) return %v, want %v", tt.tree, diameter, tt.diameter)
		}
	}
}
