package leetcode

import "testing"

func TestLeafSimilar(t *testing.T) {
	var tests = []struct {
		t1, t2  *TreeNode
		similar bool
	}{
		{newTree(1), newTree(2), false},
		{newTree(1), newTree(1, 2), false},
		{newTree(1, 2, 3), newTree(1, 2, 4), false},
		{newTree(1, 2, 3), newTree(1, 2, 3, 2, 3), false},
		{newTree(3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4), newTree(3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8), true},
	}

	for _, tt := range tests {
		similar := leafSimilar(tt.t1, tt.t2)
		if similar != tt.similar {
			t.Errorf("leafSimilar(%v, %v) return %v, want %v", tt.t1, tt.t2, similar, tt.similar)
		}
	}
}
