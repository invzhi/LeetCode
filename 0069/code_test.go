package leetcode

import "testing"

func TestMySqrt(t *testing.T) {
	var tests = []struct {
		x, sqrtx int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
		{1, 1},
	}

	for _, tt := range tests {
		sqrtx := mySqrt(tt.x)
		if sqrtx != tt.sqrtx {
			t.Errorf("mySqrt(%v) return %v, want %v", tt.x, sqrtx, tt.sqrtx)
		}
	}
}
