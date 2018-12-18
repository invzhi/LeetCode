package leetcode

import "testing"

func TestWidthOfBinaryTree(t *testing.T) {
	var tests = []struct {
		tree  *TreeNode
		width int
	}{
		{newTree(1, 3, 2, 5, 3, nil, 9), 4},
		{newTree(1, 3, nil, 5, 3), 2},
		{newTree(1, 3, 2, 5), 2},
		{newTree(1, 3, 2, 5, nil, nil, 9, 6, nil, nil, 7), 8},
	}

	for _, tt := range tests {
		width := widthOfBinaryTree(tt.tree)
		if width != tt.width {
			t.Errorf("widthOfBinaryTree(%v) return %v, want %v", tt.tree, width, tt.width)
		}
	}
}
