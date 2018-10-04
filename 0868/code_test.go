package leetcode

import "testing"

func TestBinaryGap(t *testing.T) {
	var tests = []struct {
		N   int
		gap int
	}{
		{22, 2},
		{5, 2},
		{6, 1},
		{8, 0},
	}

	for _, tt := range tests {
		gap := binaryGap(tt.N)
		if gap != tt.gap {
			t.Errorf("binaryGap(%v) return %v, want %v", tt.N, gap, tt.gap)
		}
	}
}
