package leetcode

import "testing"

func TestMaxPathSum(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		maxSum int
	}{
		{newTree(-3), -3},
		{newTree(1, 2, 3), 6},
		{newTree(-10, 9, 20, nil, nil, 15, 7), 42},
		{newTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 1), 48},
	}

	for _, tt := range tests {
		maxSum := maxPathSum(tt.tree)
		if maxSum != tt.maxSum {
			t.Errorf("maxPathSum(%v) return %v, want %v", tt.tree, maxSum, tt.maxSum)
		}
	}
}
