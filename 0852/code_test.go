package leetcode

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	var tests = []struct {
		A    []int
		peak int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
	}

	for _, tt := range tests {
		peak := peakIndexInMountainArray(tt.A)
		if peak != tt.peak {
			t.Errorf("peakIndexInMountainArray(%v) return %v, want %v", tt.A, peak, tt.peak)
		}
	}
}
