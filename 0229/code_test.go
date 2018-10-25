package leetcode

import (
	"reflect"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	var tests = []struct {
		nums []int
		elem []int
	}{
		{[]int{3, 2, 3}, []int{3}},
		{[]int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
		{[]int{1, 2, 2, 3, 2, 1, 1, 3}, []int{1, 2}},
	}

	for _, tt := range tests {
		elem := majorityElement(tt.nums)
		if reflect.DeepEqual(elem, tt.elem) == false {
			t.Errorf("majorityElement(%v) return %v, want %v", tt.nums, elem, tt.elem)
		}
	}
}
