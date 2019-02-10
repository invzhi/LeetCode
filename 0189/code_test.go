package leetcode

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		before []int
		k      int
		after  []int
	}{
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{4, 5, 6, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, 4, []int{3, 4, 5, 6, 1, 2}},
		{[]int{1, 2, 3, 4, 5, 6}, 9, []int{4, 5, 6, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for _, tt := range tests {
		before := make([]int, len(tt.before))
		copy(before, tt.before)

		rotate(tt.before, tt.k)
		if reflect.DeepEqual(tt.before, tt.after) == false {
			t.Errorf("rotate(%v, %v) return %v, want %v", before, tt.k, tt.before, tt.after)
		}
	}
}
