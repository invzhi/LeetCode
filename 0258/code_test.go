package leetcode

import "testing"

func TestAddDigits(t *testing.T) {
	var tests = []struct {
		num   int
		digit int
	}{
		{17, 8},
		{38, 2},
		{1583, 8},
		{65536, 7},
	}

	for _, tt := range tests {
		digit := addDigits(tt.num)
		if digit != tt.digit {
			t.Errorf("addDigits(%v) return %v, want %v", tt.num, digit, tt.digit)
		}
	}
}
