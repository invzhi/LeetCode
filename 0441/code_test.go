package leetcode

import "testing"

func TestArrangeCoins(t *testing.T) {
	var tests = []struct {
		n     int
		total int
	}{
		{5, 2},
		{8, 3},
		{0, 0},
		{1, 1},
	}

	for _, tt := range tests {
		total := arrangeCoins(tt.n)
		if total != tt.total {
			t.Errorf("arrangeCoins(%v) return %v, want %v", tt.n, total, tt.total)
		}
	}
}
