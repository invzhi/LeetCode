package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	var tests = []struct {
		list         []int
		reversedList *ListNode
	}{
		{
			[]int{1, 2, 3, 4, 5},
			newList(5, 4, 3, 2, 1),
		},
		{
			[]int{5, 4, 3, 2, 1},
			newList(1, 2, 3, 4, 5),
		},
	}

	for _, tt := range tests {
		reversedList := reverseList(newList(tt.list...))
		if reflect.DeepEqual(reversedList, tt.reversedList) == false {
			t.Errorf("reverseList(%v) return %v, want %v", newList(tt.list...), reversedList, tt.reversedList)
		}
	}
}
