package leetcode

import "testing"

func TestIsIsomorphic(t *testing.T) {
	var tests = []struct {
		s, t       string
		isomorphic bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"zoo", "sum", false},
		{"paper", "title", true},
		{"paper", "hello", false},
	}

	for _, tt := range tests {
		isomorphic := isIsomorphic(tt.s, tt.t)
		if isomorphic != tt.isomorphic {
			t.Errorf("isIsomorphic(%v, %v) return %v, want %v", tt.s, tt.t, isomorphic, tt.isomorphic)
		}
	}
}
