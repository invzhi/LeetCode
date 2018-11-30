package leetcode

import "testing"

func TestSumOfLeftLeaves(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		sum  int
	}{
		{newTree(), 0},
		{newTree(3, 9, 20, nil, nil, 15, 7), 24},
	}

	for _, tt := range tests {
		sum := sumOfLeftLeaves(tt.tree)
		if sum != tt.sum {
			t.Errorf("sumOfLeftLeaves(%v) return %v, want %v", tt.tree, sum, tt.sum)
		}
	}
}
