package leetcode

import (
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	var tests = []struct {
		before []int
		k      int
		after  *ListNode
	}{
		{[]int{1, 2, 3, 4, 5}, 2, newList(4, 5, 1, 2, 3)},
		{[]int{0, 1, 2}, 4, newList(2, 0, 1)},
		{[]int{0, 1, 2}, 0, newList(0, 1, 2)},
		{[]int{}, 0, newList()},
	}

	for _, tt := range tests {
		after := rotateRight(newList(tt.before...), tt.k)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("rotateRight(%v, %v) return %v, want %v", newList(tt.before...), tt.k, after, tt.after)
		}
	}
}
