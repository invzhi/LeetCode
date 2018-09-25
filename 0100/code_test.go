package leetcode

import "testing"

func TestIsSameTree(t *testing.T) {
	var tests = []struct {
		p, q   *TreeNode
		isSame bool
	}{
		{newTree(1, 2, 3), newTree(1, 2, 3), true},
		{newTree(1, 2), newTree(1, nil, 2), false},
		{newTree(1, 2, 1), newTree(1, 1, 2), false},
		{newTree(10, 5, 15), newTree(10, 5, nil, nil, 15), false},
	}

	for _, tt := range tests {
		isSame := isSameTree(tt.p, tt.q)
		if isSame != tt.isSame {
			t.Errorf("isSameTree(%v, %v) return %v, want %v", tt.p, tt.q, isSame, tt.isSame)
		}
	}
}
