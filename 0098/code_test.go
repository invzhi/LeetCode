package leetcode

import "testing"

func TestIsValidBST(t *testing.T) {
	var tests = []struct {
		tree     *TreeNode
		validBST bool
	}{
		{newTree(), true},
		{newTree(2, 1, 3), true},
		{newTree(4, 2, 6, 1, 5), false},
		{newTree(5, 1, 4, nil, nil, 3, 6), false},
		{newTree(3, nil, 30, 10, nil, nil, 15, nil, 45), false},
	}

	for _, tt := range tests {
		validBST := isValidBST(tt.tree)
		if validBST != tt.validBST {
			t.Errorf("isValidBST(%v) return %v, want %v", tt.tree, validBST, tt.validBST)
		}
	}
}
