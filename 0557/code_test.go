package leetcode

import "testing"

func TestReverseWords(t *testing.T) {
	var tests = []struct {
		s string
		r string
	}{
		{" Hello  World! ", " olleH  !dlroW "},
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}

	for _, tt := range tests {
		r := reverseWords(tt.s)
		if r != tt.r {
			t.Errorf("reverseWords(%v) return %v, want %v", tt.s, r, tt.r)
		}
	}
}
