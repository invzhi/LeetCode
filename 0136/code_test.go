package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	var tests = []struct {
		nums   []int
		single int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{nil, 0},
	}

	for _, tt := range tests {
		single := singleNumber(tt.nums)
		if single != tt.single {
			t.Errorf("singleNumber(%v) return %v, want %v", tt.nums, single, tt.single)
		}
	}
}
