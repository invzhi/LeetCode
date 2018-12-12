package leetcode

import "testing"

func TestIsSubtree(t *testing.T) {
	var tests = []struct {
		s, t *TreeNode
		sub  bool
	}{
		{newTree(3, 4, 5, 1, 2), newTree(4, 1, 2), true},
		{newTree(3, 4, 5, 1, 2, nil, nil, nil, nil, 0), newTree(4, 1, 2), false},
	}

	for _, tt := range tests {
		sub := isSubtree(tt.s, tt.t)
		if sub != tt.sub {
			t.Errorf("isSubtree(%v, %v) return %v, want %v", tt.s, tt.t, sub, tt.sub)
		}
	}
}
