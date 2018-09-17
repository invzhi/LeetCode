package leetcode

import "testing"

func TestLengthOfSubstring(t *testing.T) {
	var tests = []struct {
		s      string
		length int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tt := range tests {
		length := lengthOfLongestSubstring(tt.s)
		if length != tt.length {
			t.Errorf("lengthOfLongestSubstring(%v) return %v, want %v", tt.s, length, tt.length)
		}
	}
}
