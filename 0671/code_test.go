package leetcode

import "testing"

func TestFindSecondMinimumValue(t *testing.T) {
	var tests = []struct {
		tree  *TreeNode
		value int
	}{
		{newTree(2, 2, 5, nil, nil, 5, 7), 5},
		{newTree(2, 2, 2), -1},
	}

	for _, tt := range tests {
		value := findSecondMinimumValue(tt.tree)
		if value != tt.value {
			t.Errorf("findSecondMinimumValue(%v) return %v, want %v", tt.tree, value, tt.value)
		}
	}
}
