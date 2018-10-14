package leetcode

import "testing"

func TestReverseVowels(t *testing.T) {
	var tests = []struct {
		s string
		r string
	}{
		{"aA", "Aa"},
		{"hello", "holle"},
		{"leetcode", "leotcede"},
	}

	for _, tt := range tests {
		r := reverseVowels(tt.s)
		if r != tt.r {
			t.Errorf("reverseVowels(%v) return %v, want %v", tt.s, r, tt.r)
		}
	}
}
