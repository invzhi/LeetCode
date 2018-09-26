package leetcode

import "testing"

func TestList(t *testing.T) {
	var tests = []struct {
		vals []int
		s    string
	}{
		{[]int{}, "<nil>"},
		{[]int{1, 2, 3, 4}, "(1 -> 2 -> 3 -> 4)"},
	}

	for _, tt := range tests {
		s := newList(tt.vals...).String()
		if s != tt.s {
			t.Errorf("newList(%v) return %v, want %v", tt.vals, s, tt.s)
		}
	}
}
