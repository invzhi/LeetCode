package leetcode

import "testing"

func TestFindKthLargest(t *testing.T) {
	var tests = []struct {
		nums    []int
		k       int
		element int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 1, 6},
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 1, 5, 6, 4}, 3, 4},
		{[]int{3, 2, 1, 5, 6, 4}, 4, 3},
		{[]int{3, 2, 1, 5, 6, 4}, 5, 2},
		{[]int{3, 2, 1, 5, 6, 4}, 6, 1},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)

		element := findKthLargest(tt.nums, tt.k)
		if element != tt.element {
			t.Errorf("findKthLargest(%v, %v) return %v, want %v", nums, tt.k, element, tt.element)
		}
	}
}
