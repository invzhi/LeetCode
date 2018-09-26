package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	var tests = []struct {
		nums   []int
		single int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	}

	for _, tt := range tests {
		single := singleNumber(tt.nums)
		if single != tt.single {
			t.Errorf("singleNumber(%v) return %v, want %v", tt.nums, single, tt.single)
		}
	}
}
