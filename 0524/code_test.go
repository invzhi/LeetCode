package leetcode

import "testing"

func TestFindLongestWord(t *testing.T) {
	var tests = []struct {
		s string
		d []string
		w string
	}{
		{"abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
		{"abpcplea", []string{"a", "b", "c"}, "a"},
	}

	for _, tt := range tests {
		w := findLongestWord(tt.s, tt.d)
		if w != tt.w {
			t.Errorf("findLongestWord(%v, %v) return %v, want %v", tt.s, tt.d, w, tt.w)
		}
	}
}
