package leetcode

import "testing"

func TestReverseString(t *testing.T) {
	var tests = []struct {
		s string
		r string
	}{
		{"hello", "olleh"},
		{"A man, a plan, a canal: Panama", "amanaP :lanac a ,nalp a ,nam A"},
	}

	for _, tt := range tests {
		r := reverseString(tt.s)
		if r != tt.r {
			t.Errorf("reverseString(%v) return %v, want %v", tt.s, r, tt.r)
		}
	}
}
