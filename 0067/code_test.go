package leetcode

import "testing"

func TestAddBinary(t *testing.T) {
	var tests = []struct {
		a, b string
		sum  string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, tt := range tests {
		sum := addBinary(tt.a, tt.b)
		if sum != tt.sum {
			t.Errorf("addBinary(%v, %v) return %v, want %v", tt.a, tt.b, sum, tt.sum)
		}
	}
}
