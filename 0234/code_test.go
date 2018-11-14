package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		list       *ListNode
		palindrome bool
	}{
		{newList(1, 2), false},
		{newList(1, 2, 2, 1), true},
		{newList(1, 2, 3, 2, 1), true},
	}

	for _, tt := range tests {
		palindrome := isPalindrome(tt.list)
		if palindrome != tt.palindrome {
			t.Errorf("isPalindrome(%v) return %v, want %v", tt.list, palindrome, tt.palindrome)
		}
	}
}
