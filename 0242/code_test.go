package leetcode

import "testing"

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		s, t      string
		isAnagram bool
	}{
		{
			"anagram", "nagaram",
			true,
		},
		{
			"rat", "car",
			false,
		},
	}

	for _, tt := range tests {
		anagram := isAnagram(tt.s, tt.t)
		if anagram != tt.isAnagram {
			t.Errorf("isAnagram(%v, %v) return %v, want %v", tt.s, tt.t, anagram, tt.isAnagram)
		}
	}
}
