package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		numbers []int
		target  int
		indices []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 5, 6, 7}, 12, []int{2, 4}},
		{[]int{2, 6, 8, 23}, 16, nil},
	}

	for _, tt := range tests {
		indices := twoSum(tt.numbers, tt.target)
		if reflect.DeepEqual(indices, tt.indices) == false {
			t.Errorf("twoSum(%v, %v) return %v, want %v", tt.numbers, tt.target, indices, tt.indices)
		}
	}
}
