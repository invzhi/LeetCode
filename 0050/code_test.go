package leetcode

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	var tests = []struct {
		x  float64
		n  int
		xn float64
	}{
		{x: 2, n: 10},
		{x: 2.1, n: 3},
		{x: 2, n: -2},
		{x: 1, n: math.MaxInt32},
		{x: -2.1, n: math.MinInt32},
	}

	for _, tt := range tests {
		tt.xn = math.Pow(tt.x, float64(tt.n))

		xn := myPow(tt.x, tt.n)
		if xn != tt.xn {
			t.Errorf("myPow(%v, %v) return %v, want %v", tt.x, tt.n, xn, tt.xn)
		}
	}
}
