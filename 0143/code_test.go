package leetcode

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	var tests = []struct {
		before []int
		after  *ListNode
	}{
		{[]int{1, 2, 3, 4}, newList(1, 4, 2, 3)},
		{[]int{1, 2, 3, 4, 5}, newList(1, 5, 2, 4, 3)},
	}

	for _, tt := range tests {
		list := newList(tt.before...)
		reorderList(list)
		if reflect.DeepEqual(list, tt.after) == false {
			t.Errorf("reorderList(%v) return %v, want %v", newList(tt.before...), list, tt.after)
		}
	}
}
