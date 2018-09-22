package leetcode

import "testing"

func TestCountAndSay(t *testing.T) {
	var tests = []struct {
		n int
		s string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
	}

	for _, tt := range tests {
		s := countAndSay(tt.n)
		if s != tt.s {
			t.Errorf("countAndSay(%v) return %v, want %v", tt.n, s, tt.s)
		}
	}
}
