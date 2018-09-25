package leetcode

import "testing"

func TestSumNumbers(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		sum  int
	}{
		{newTree(1, 2, 3), 25},
		{newTree(1, 2, 2, 3), 135},
		{newTree(4, 9, 0, 5, 1), 1026},
	}

	for _, tt := range tests {
		sum := sumNumbers(tt.tree)
		if sum != tt.sum {
			t.Errorf("sumNumbers(%v) return %v, want %v", tt.tree, sum, tt.sum)
		}
	}
}
