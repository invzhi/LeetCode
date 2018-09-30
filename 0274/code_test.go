package leetcode

import "testing"

func TestHIndex(t *testing.T) {
	var tests = []struct {
		citations []int
		h         int
	}{
		{nil, 0},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{3, 0, 6, 1, 5}, 3},
	}

	for _, tt := range tests {
		h := hIndex(tt.citations)
		if h != tt.h {
			t.Errorf("hIndex(%v) return %v, want %v", tt.citations, h, tt.h)
		}
	}
}
