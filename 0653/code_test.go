package leetcode

import "testing"

func TestFindTarget(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		target int
		exist  bool
	}{
		{newTree(5, 3, 6, 2, 4, nil, 7), 9, true},
		{newTree(5, 3, 6, 2, 4, nil, 7), 4, false},
		{newTree(5, 3), 8, true},
		{nil, 2, false},
	}

	for _, tt := range tests {
		exist := findTarget(tt.tree, tt.target)
		if exist != tt.exist {
			t.Errorf("findTarget(%v, %v) return %v, want %v", tt.tree, tt.target, exist, tt.exist)
		}
	}
}
