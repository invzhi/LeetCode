package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		strs   []string
		prefix string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"flower", "flow"}, "flow"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"aa", "a"}, "a"},
		{[]string{}, ""},
	}

	for _, tt := range tests {
		prefix := longestCommonPrefix(tt.strs)
		if prefix != tt.prefix {
			t.Errorf("longestCommonPrefix(%q) return %q, want %q", tt.strs, prefix, tt.prefix)
		}
	}
}
