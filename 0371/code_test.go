package leetcode

import "testing"

func TestGetSum(t *testing.T) {
	var tests = []struct {
		a, b int
	}{
		{1, 2},
		{-2, 3},
		{-1, -1},
	}

	for _, tt := range tests {
		sum := getSum(tt.a, tt.b)
		if tt.a+tt.b != sum {
			t.Errorf("getSum(%v, %v) return %v, want %v", tt.a, tt.b, sum, tt.a+tt.b)
		}
	}
}
