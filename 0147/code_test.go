package leetcode

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func newList(vs ...int) *ListNode {
	l := new(ListNode)
	n := l
	for _, v := range vs {
		n.Next = &ListNode{Val: v}
		n = n.Next
	}
	return l.Next
}

func (l *ListNode) String() string {
	var b strings.Builder
	b.WriteByte('(')
	if l != nil {
		b.WriteString(strconv.Itoa(l.Val))
		l = l.Next
	}
	for l != nil {
		b.WriteString(" -> ")
		b.WriteString(strconv.Itoa(l.Val))
		l = l.Next
	}
	b.WriteByte(')')
	return b.String()
}

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
