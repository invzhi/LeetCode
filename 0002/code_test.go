package leetcode

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		l1, l2 *ListNode
		sum    *ListNode
	}{
		{newList(2, 4, 6), newList(5, 6, 4), newList(7, 0, 1, 1)},
		{newList(2, 4, 3), newList(5, 6, 4), newList(7, 0, 8)},
		{newList(2, 4), newList(5, 6, 4), newList(7, 0, 5)},
		{newList(2), newList(5, 6, 4), newList(7, 6, 4)},
		{newList(2), newList(5, 6), newList(7, 6)},
		{newList(2), newList(5), newList(7)},
	}

	for _, tt := range tests {
		sum := addTwoNumbers(tt.l1, tt.l2)
		if reflect.DeepEqual(sum, tt.sum) == false {
			t.Errorf("addTwoNumbers(%v, %v) return %v, want %v", tt.l1, tt.l2, sum, tt.sum)
		}
	}
}
