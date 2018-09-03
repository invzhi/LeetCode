package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	var tests = []struct {
		nums []int
		once []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{[]int{3, 2, 3, 7, 2, 1}, []int{1, 7}},
		{[]int{2, 1, 2, 3, 4, 1}, []int{3, 4}},
		{[]int{1, 2}, []int{1, 2}},
	}

	for _, tt := range tests {
		once := singleNumber(tt.nums)
		if len(once) != 2 {
			t.Fatalf("singleNumber(%v) return %v, len(%[2]v) != 2", tt.nums, once)
		}
		if (once[0] != tt.once[0] || once[1] != tt.once[1]) && (once[0] != tt.once[1] || once[1] != tt.once[0]) {
			t.Errorf("singleNumber(%v) return %v, want %v", tt.nums, once, tt.once)
		}
	}
}
