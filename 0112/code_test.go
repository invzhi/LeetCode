package leetcode

import "testing"

func TestHasPathSum(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		sum    int
		hasSum bool
	}{
		{newTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1), 22, true},
		{newTree(5, 4, 8, 11, nil, 13, 4, nil, 2, nil, nil, nil, 1), 22, true},
	}

	for _, tt := range tests {
		hasSum := hasPathSum(tt.tree, tt.sum)
		if hasSum != tt.hasSum {
			t.Errorf("hasPathSum(%v, %v) return %v, want %v", tt.tree, tt.sum, hasSum, tt.hasSum)
		}
	}
}
