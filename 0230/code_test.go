package leetcode

import "testing"

func TestKthSmallest(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		k    int
		kth  int
	}{
		{newTree(3, 1, 4, nil, 2), 1, 1},
		{newTree(3, 1, 4, nil, 2), 2, 2},
		{newTree(3, 1, 4, nil, 2), 3, 3},
		{newTree(3, 1, 4, nil, 2), 4, 4},
		{newTree(5, 3, 6, 2, 4, nil, nil, 1), 1, 1},
		{newTree(5, 3, 6, 2, 4, nil, nil, 1), 2, 2},
		{newTree(5, 3, 6, 2, 4, nil, nil, 1), 3, 3},
		{newTree(5, 3, 6, 2, 4, nil, nil, 1), 4, 4},
		{newTree(5, 3, 6, 2, 4, nil, nil, 1), 5, 5},
		{newTree(5, 3, 6, 2, 4, nil, nil, 1), 6, 6},
	}

	for _, tt := range tests {
		kth := kthSmallest(tt.tree, tt.k)
		if kth != tt.kth {
			t.Errorf("kthSmallest(%v, %v) return %v, want %v", tt.tree, tt.k, kth, tt.kth)
		}
	}
}
