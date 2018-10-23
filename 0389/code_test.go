package leetcode

import "testing"

func TestFindTheDifference(t *testing.T) {
	var tests = []struct {
		s, t   string
		letter byte
	}{
		{"abcd", "abcde", 'e'},
		{"abcdd", "dcdbae", 'e'},
	}

	for _, tt := range tests {
		letter := findTheDifference(tt.s, tt.t)
		if letter != tt.letter {
			t.Errorf("findTheDifference(%v, %v) return %v, want %v", tt.s, tt.t, letter, tt.letter)
		}
	}
}
