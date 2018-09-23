package leetcode

import "testing"

func TestSmallestRangeI(t *testing.T) {
	var tests = []struct {
		A []int
		K int
		D int
	}{
		{[]int{1}, 0, 0},
		{[]int{0, 10}, 2, 6},
		{[]int{10, 0}, 2, 6},
		{[]int{1, 3, 6}, 3, 0},
	}

	for _, tt := range tests {
		D := smallestRangeI(tt.A, tt.K)
		if D != tt.D {
			t.Errorf("smallestRangeI(%v, %v) return %v, want %v", tt.A, tt.K, D, tt.D)
		}
	}
}
