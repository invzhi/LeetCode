package leetcode

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		integer  int
		reversed int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{math.MaxInt32, 0},
		{math.MinInt32, 0},
	}

	for _, tt := range tests {
		reversed := reverse(tt.integer)
		if reversed != tt.reversed {
			t.Errorf("reverse(%v) return %v, want %v", tt.integer, reversed, tt.reversed)
		}
	}
}
