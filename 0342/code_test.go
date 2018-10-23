package leetcode

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	var tests = []struct {
		n int
		b bool
	}{
		{0, false},
		{5, false},
		{16, true},
	}

	for _, tt := range tests {
		b := isPowerOfFour(tt.n)
		if b != tt.b {
			t.Errorf("isPowerOfFour(%v) return %v, want %v", tt.n, b, tt.b)
		}
	}
}
