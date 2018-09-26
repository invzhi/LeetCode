package leetcode

import (
	"reflect"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	var tests = []struct {
		list       []int
		sortedList *ListNode
	}{
		{
			[]int{4, 2, 1, 3},
			newList(1, 2, 3, 4),
		},
		{
			[]int{-1, 5, 3, 4, 0},
			newList(-1, 0, 3, 4, 5),
		},
		{
			[]int{1, 2, 3, 4},
			newList(1, 2, 3, 4),
		},
		{
			[]int{4, 3, 2, 1},
			newList(1, 2, 3, 4),
		},
		{
			nil,
			nil,
		},
	}

	for _, tt := range tests {
		sortedList := insertionSortList(newList(tt.list...))
		if reflect.DeepEqual(sortedList, tt.sortedList) == false {
			t.Errorf("insertionSortList(%v) return %v, want %v", newList(tt.list...), sortedList, tt.sortedList)
		}
	}
}
