package leetcode

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	var tests = []struct {
		n int
		b bool
	}{
		{0, false},
		{9, true},
		{27, true},
		{45, false},
	}

	for _, tt := range tests {
		b := isPowerOfThree(tt.n)
		if b != tt.b {
			t.Errorf("isPowerOfThree(%v) return %v, want %v", tt.n, b, tt.b)
		}
	}
}
