package leetcode

import "testing"

func TestLargestNumber(t *testing.T) {
	var tests = []struct {
		nums []int
		s    string
	}{
		{[]int{0, 0}, "0"},
		{[]int{10, 2}, "210"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
	}

	for _, tt := range tests {
		s := largestNumber(tt.nums)
		if s != tt.s {
			t.Errorf("largestNumber(%v) return %v, want %v", tt.nums, s, tt.s)
		}
	}
}
