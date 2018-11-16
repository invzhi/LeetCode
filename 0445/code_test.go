package leetcode

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		l1, l2 *ListNode
		l      *ListNode
	}{
		{newList(0), newList(0), newList(0)},
		{newList(7), newList(0), newList(7)},
		{newList(5), newList(5), newList(1, 0)},
		{newList(1), newList(9, 9), newList(1, 0, 0)},
		{newList(1, 2), newList(3, 8, 8), newList(4, 0, 0)},
		{newList(7, 2, 4, 3), newList(0), newList(7, 2, 4, 3)},
		{newList(7, 2, 4, 3), newList(5, 6, 4), newList(7, 8, 0, 7)},
		{newList(7, 2, 4, 3), newList(7, 5, 7), newList(8, 0, 0, 0)},
	}

	for _, tt := range tests {
		l := addTwoNumbers(tt.l1, tt.l2)
		if reflect.DeepEqual(l, tt.l) == false {
			t.Errorf("addTwoNumbers(%v, %v) return %v, want %v", tt.l1, tt.l2, l, tt.l)
		}
	}
}
