package leetcode

import "testing"

func TestGetMinimumDifference(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		diff int
	}{
		{newTree(1, nil, 3, 2), 1},
		{newTree(10, 4, 15, 1, 7, 13, 18), 2},
	}

	for _, tt := range tests {
		diff := getMinimumDifference(tt.tree)
		if diff != tt.diff {
			t.Errorf("getMinimumDifference(%v) return %v, want %v", tt.tree, diff, tt.diff)
		}
	}
}
