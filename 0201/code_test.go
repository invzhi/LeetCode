package leetcode

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
	var tests = []struct {
		m, n int
		and  int
	}{
		{0, 1, 0},
		{1, 1, 1},
		{5, 7, 4},
		{41, 44, 40},
	}

	for _, tt := range tests {
		and := rangeBitwiseAnd(tt.m, tt.n)
		if and != tt.and {
			t.Errorf("rangeBitwiseAnd(%v, %v) return %v, want %v", tt.m, tt.n, and, tt.and)
		}
	}
}
