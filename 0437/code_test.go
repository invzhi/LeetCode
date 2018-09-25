package leetcode

import "testing"

func TestPathSum(t *testing.T) {
	var tests = []struct {
		tree    *TreeNode
		sum     int
		numPath int
	}{
		{newTree(10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1), 8, 3},
		{newTree(1, -2, -3, 1, 3, -2, nil, -1), 1, 3},
	}

	for _, tt := range tests {
		numPath := pathSum(tt.tree, tt.sum)
		if numPath != tt.numPath {
			t.Errorf("PathSum(%v, %v) return %v, want %v", tt.tree, tt.sum, numPath, tt.numPath)
		}
	}
}
