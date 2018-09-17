package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	var tests = []struct {
		nums1, nums2 []int
		median       float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 3}, []int{2, 4}, 2.5},
		{[]int{1, 4}, []int{2, 3}, 2.5},
		{[]int{2, 3}, []int{1, 4}, 2.5},
		{[]int{2, 4}, []int{1, 3}, 2.5},
		{[]int{3, 4}, []int{1, 2}, 2.5},
	}

	for _, tt := range tests {
		median := findMedianSortedArrays(tt.nums1, tt.nums2)
		if median != tt.median {
			t.Errorf("findMedianSortedArrays(%v, %v) return %v, want %v", tt.nums1, tt.nums2, median, tt.median)
		}
	}
}
