package leetcode

import "testing"

func TestNumTrees(t *testing.T) {
	var tests = []struct {
		n   int
		num int
	}{
		{3, 5},
		{10, 16796},
	}

	for _, tt := range tests {
		num := numTrees(tt.n)
		if num != tt.num {
			t.Errorf("numTrees(%v) return %v, want %v", tt.n, num, tt.num)
		}
	}
}
