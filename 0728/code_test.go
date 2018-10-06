package leetcode

import (
	"reflect"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	var tests = []struct {
		left, right int
		nums        []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
	}

	for _, tt := range tests {
		nums := selfDividingNumbers(tt.left, tt.right)
		if reflect.DeepEqual(nums, tt.nums) == false {
			t.Errorf("selfDividingNumbers(%v, %v) return %v, want %v", tt.left, tt.right, nums, tt.nums)
		}
	}
}
