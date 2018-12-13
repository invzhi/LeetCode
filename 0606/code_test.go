package leetcode

import "testing"

func TestTree2str(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		str  string
	}{
		{newTree(), ""},
		{newTree(1, 2, 3, 4), "1(2(4))(3)"},
		{newTree(1, 2, 3, nil, 4), "1(2()(4))(3)"},
	}

	for _, tt := range tests {
		str := tree2str(tt.tree)
		if str != tt.str {
			t.Errorf("tree2str(%v) return %v, want %v", tt.tree, str, tt.str)
		}
	}
}
