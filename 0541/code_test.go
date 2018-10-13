package leetcode

import "testing"

func TestReverseStr(t *testing.T) {
	var tests = []struct {
		s string
		k int
		r string
	}{
		{"abcdefg", 2, "bacdfeg"},
		{"abcdefgh", 3, "cbadefhg"},
	}

	for _, tt := range tests {
		r := reverseStr(tt.s, tt.k)
		if r != tt.r {
			t.Errorf("reverseStr(%v, %v) return %v, want %v", tt.s, tt.k, r, tt.r)
		}
	}
}
