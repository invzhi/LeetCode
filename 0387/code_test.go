package leetcode

import "testing"

func TestFirstUniqChar(t *testing.T) {
	var tests = []struct {
		s string
		i int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"leetcodeleetcode", -1},
	}

	for _, tt := range tests {
		i := firstUniqChar(tt.s)
		if i != tt.i {
			t.Errorf("firstUniqChar(%v) return %v, want %v", tt.s, i, tt.i)
		}
	}
}
