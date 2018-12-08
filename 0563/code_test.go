package leetcode

import "testing"

func TestFindTilt(t *testing.T) {
	var tests = []struct {
		tree *TreeNode
		tilt int
	}{
		{newTree(1, 2, 3), 1},
	}

	for _, tt := range tests {
		tilt := findTilt(tt.tree)
		if tilt != tt.tilt {
			t.Errorf("findTilt(%v) return %v, want %v", tt.tree, tilt, tt.tilt)
		}
	}
}
