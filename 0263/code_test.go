package leetcode

import "testing"

func TestIsUgly(t *testing.T) {
	var tests = []struct {
		num  int
		ugly bool
	}{
		{0, false},
		{6, true},
		{8, true},
		{14, false},
	}

	for _, tt := range tests {
		ugly := isUgly(tt.num)
		if ugly != tt.ugly {
			t.Errorf("isUgly(%v) return %v, want %v", tt.num, ugly, tt.ugly)
		}
	}
}
