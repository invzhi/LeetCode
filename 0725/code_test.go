package leetcode

import (
	"reflect"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	var tests = []struct {
		list  []int
		k     int
		parts []*ListNode
	}{
		{[]int{1, 2, 3}, 5, []*ListNode{newList(1), newList(2), newList(3), nil, nil}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, []*ListNode{newList(1, 2, 3, 4), newList(5, 6, 7), newList(8, 9, 10)}},
	}

	for _, tt := range tests {
		parts := splitListToParts(newList(tt.list...), tt.k)
		if reflect.DeepEqual(parts, tt.parts) == false {
			t.Errorf("splitListToParts(%v, %v) return %v, want %v", newList(tt.list...), tt.k, parts, tt.parts)
		}
	}
}
