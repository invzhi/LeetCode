package leetcode

import "testing"

func TestHasGroupsSizeX(t *testing.T) {
	var tests = []struct {
		deck []int
		has  bool
	}{
		{[]int{1}, false},
		{[]int{1, 1}, true},
		{[]int{1, 1, 1, 2, 2, 2}, true},
		{[]int{1, 1, 2, 2, 2, 2}, true},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, true},
		{[]int{1, 1, 1, 2, 2, 2, 3, 3}, false},
		{[]int{0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3}, true},
	}

	for _, tt := range tests {
		has := hasGroupsSizeX(tt.deck)
		if has != tt.has {
			t.Errorf("hasGroupsSizeX(%v) return %v, want %v", tt.deck, has, tt.has)
		}
	}
}
