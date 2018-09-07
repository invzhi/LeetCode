package leetcode

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	var tests = []struct {
		num       int
		isPerfect bool
	}{
		{16, true},
		{14, false},
		{0, true},
	}

	for _, tt := range tests {
		isPerfect := isPerfectSquare(tt.num)
		if isPerfect != tt.isPerfect {
			t.Errorf("isPerfectSquare(%v) return %v, want %v", tt.num, isPerfect, tt.isPerfect)
		}
	}
}
