package leetcode

import "testing"

func TestCheckPerfectNumber(t *testing.T) {
	var tests = []struct {
		num     int
		perfect bool
	}{
		{-1, false},
		{0, false},
		{1, false},
		{3, false},
		{6, true},
		{28, true},
		{496, true},
		{8128, true},
		{33550336, true},
	}

	for _, tt := range tests {
		perfect := checkPerfectNumber(tt.num)
		if perfect != tt.perfect {
			t.Errorf("checkPerfectNumber(%v) return %v, want %v", tt.num, perfect, tt.perfect)
		}
	}
}
