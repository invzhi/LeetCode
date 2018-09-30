package leetcode

import "testing"

func TestHIndex(t *testing.T) {
	var tests = []struct {
		citations []int
		h         int
	}{
		{nil, 0},
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{1, 1}, 1},
		{[]int{11, 15}, 2},
		{[]int{0, 1, 3, 5, 6}, 3},
	}

	for _, tt := range tests {
		h := hIndex(tt.citations)
		if h != tt.h {
			t.Errorf("hIndex(%v) return %v, want %v", tt.citations, h, tt.h)
		}
	}
}
