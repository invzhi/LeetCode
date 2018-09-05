package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	var tests = []struct {
		nums1, nums2 []int
		intersection []int
	}{
		{
			[]int{1, 2, 2, 1}, []int{2, 2},
			[]int{2, 2},
		},
		{
			[]int{4, 9, 5}, []int{9, 4, 9, 8, 4},
			[]int{4, 9},
		},
	}

	for _, tt := range tests {
		nums := intersect(tt.nums1, tt.nums2)
		sort.Ints(nums)
		if reflect.DeepEqual(nums, tt.intersection) == false {
			t.Errorf("intersect(%v, %v) return %v, want %v", tt.nums1, tt.nums2, nums, tt.intersection)
		}
	}
}
