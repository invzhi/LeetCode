package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var tests = []struct {
		l1, l2 []int
		l      *ListNode
	}{
		{
			[]int{1, 2, 4}, []int{1, 3, 4},
			newList(1, 1, 2, 3, 4, 4),
		},
		{
			[]int{1, 2, 3}, []int{4, 5, 6},
			newList(1, 2, 3, 4, 5, 6),
		},
		{
			[]int{2, 4, 6}, []int{1, 3, 5},
			newList(1, 2, 3, 4, 5, 6),
		},
	}

	for _, tt := range tests {
		l := mergeTwoLists(newList(tt.l1...), newList(tt.l2...))
		if reflect.DeepEqual(l, tt.l) == false {
			t.Errorf("mergeTwoLists(%v, %v) return %v, want %v", newList(tt.l1...), newList(tt.l2...), l, tt.l)
		}
	}
}
