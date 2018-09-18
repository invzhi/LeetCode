package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		integer    int
		palindrome bool
	}{
		{1221, true},
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, tt := range tests {
		palindrome := isPalindrome(tt.integer)
		if palindrome != tt.palindrome {
			t.Errorf("isPalindrome(%v) return %v, want %v", tt.integer, palindrome, tt.palindrome)
		}
	}
}
