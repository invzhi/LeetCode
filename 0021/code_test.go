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
