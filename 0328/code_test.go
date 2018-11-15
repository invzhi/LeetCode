package leetcode

import (
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	var tests = []struct {
		before []int
		after  *ListNode
	}{
		{[]int{}, nil},
		{[]int{1}, newList(1)},
		{[]int{1, 2, 3, 4, 5}, newList(1, 3, 5, 2, 4)},
		{[]int{1, 2, 3, 4, 5, 6}, newList(1, 3, 5, 2, 4, 6)},
		{[]int{2, 1, 3, 5, 6, 4, 7}, newList(2, 3, 6, 7, 1, 5, 4)},
	}

	for _, tt := range tests {
		after := oddEvenList(newList(tt.before...))
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("oddEvenList(%v) return %v, want %v", newList(tt.before...), after, tt.after)
		}
	}
}
