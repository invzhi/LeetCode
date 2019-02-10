package leetcode

import "testing"

func TestCountPrimes(t *testing.T) {
	var tests = []struct {
		n   int
		cnt int
	}{
		{0, 0},
		{10, 4},
		{100, 25},
		{1000, 168},
		{10000, 1229},
	}

	for _, tt := range tests {
		cnt := countPrimes(tt.n)
		if cnt != tt.cnt {
			t.Errorf("countPrimes(%v) return %v, want %v", tt.n, cnt, tt.cnt)
		}
	}
}
