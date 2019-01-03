package leetcode

import (
	"reflect"
	"testing"
)

func findNode(l *ListNode, v int) *ListNode {
	for l != nil {
		if l.Val == v {
			return l
		}
		l = l.Next
	}
	return nil
}

func TestDeleteNode(t *testing.T) {
	var tests = []struct {
		before []int
		v      int
		after  *ListNode
	}{
		{[]int{4, 5, 1, 9}, 5, newList(4, 1, 9)},
		{[]int{4, 5, 1, 9}, 1, newList(4, 5, 9)},
		{[]int{1, 2, 3, 4, 5}, 3, newList(1, 2, 4, 5)},
	}

	for _, tt := range tests {
		before := newList(tt.before...)

		deleteNode(findNode(before, tt.v))
		if reflect.DeepEqual(before, tt.after) == false {
			t.Errorf("deleteNode(%v, %v) return %v, want %v", newList(tt.before...), tt.v, before, tt.after)
		}
	}
}
