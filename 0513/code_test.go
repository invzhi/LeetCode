package leetcode

import "testing"

func TestFindBottomLeftValue(t *testing.T) {
	var tests = []struct {
		tree  *TreeNode
		value int
	}{
		{newTree(1), 1},
		{newTree(2, 1, 3), 1},
		{newTree(1, 2, 3, 4, nil, 5, 6, nil, nil, 7), 7},
	}

	for _, tt := range tests {
		value := findBottomLeftValue(tt.tree)
		if value != tt.value {
			t.Errorf("findBottomLeftValue(%v) return %v, want %v", tt.tree, value, tt.value)
		}
	}
}
