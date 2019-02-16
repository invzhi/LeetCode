package leetcode

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	var tests = []struct {
		nums     []int
		products []int
	}{
		{[]int{0, 0}, []int{0, 0}},
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	}

	for _, tt := range tests {
		products := productExceptSelf(tt.nums)
		if reflect.DeepEqual(products, tt.products) == false {
			t.Errorf("productExceptSelf(%v) return %v, want %v", tt.nums, products, tt.products)
		}
	}
}
