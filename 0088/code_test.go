package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		nums  []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0}, 3,
			[]int{2, 5, 6}, 3,
			[]int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, tt := range tests {
		nums1 := make([]int, len(tt.nums1))
		copy(nums1, tt.nums1)

		merge(tt.nums1, tt.m, tt.nums2, tt.n)
		if reflect.DeepEqual(tt.nums1, tt.nums) == false {
			t.Errorf("merge(%v, %v, %v, %v) return %v, want %v", nums1, tt.m, tt.nums2, tt.n, tt.nums1, tt.nums)
		}
	}
}

func TestPosition(t *testing.T) {
	var tests = []struct {
		nums   []int
		num    int
		lo, hi int
		i      int
	}{
		{
			[]int{1, 3, 3, 0, 0, 0}, 2,
			0, 2,
			1,
		},
		{
			[]int{1, 2, 3, 0, 0, 0}, 2,
			0, 2,
			2,
		},
		{
			[]int{1, 2, 2, 3, 0, 0}, 5,
			3, 3,
			4,
		},
		{
			[]int{1, 2, 2, 3, 5, 0}, 6,
			5, 4,
			5,
		},
	}

	for _, tt := range tests {
		i := position(tt.nums, tt.num, tt.lo, tt.hi)
		if i != tt.i {
			t.Errorf("position(%v, %v, %v, %v) return %v, want %v", tt.nums, tt.num, tt.lo, tt.hi, i, tt.i)
		}
	}
}
