package leetcode

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	var tests = []struct {
		n int
		b bool
	}{
		{1, true},
		{16, true},
		{218, false},
	}

	for _, tt := range tests {
		b := isPowerOfTwo(tt.n)
		if b != tt.b {
			t.Errorf("isPowerOfTwo(%v) return %v, want %v", tt.n, b, tt.b)
		}
	}
}
