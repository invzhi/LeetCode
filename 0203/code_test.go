package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	var tests = []struct {
		list    []int
		val     int
		removed *ListNode
	}{
		{[]int{1, 1}, 1, newList()},
		{[]int{1, 2, 6, 3, 4, 5, 6}, 6, newList(1, 2, 3, 4, 5)},
	}

	for _, tt := range tests {
		removed := removeElements(newList(tt.list...), tt.val)
		if reflect.DeepEqual(removed, tt.removed) == false {
			t.Errorf("removeElements(%v, %v) return %v, want %v", newList(tt.list...), tt.val, removed, tt.removed)
		}
	}
}
