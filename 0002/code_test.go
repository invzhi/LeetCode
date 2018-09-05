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
		{newList(), newList(5), newList(5)},
		{newList(), newList(), newList()},
	}

	for _, tt := range tests {
		sum := addTwoNumbers(tt.l1, tt.l2)
		if reflect.DeepEqual(sum, tt.sum) == false {
			t.Errorf("addTwoNumbers(%v, %v) return %v, want %v", tt.l1, tt.l2, sum, tt.sum)
		}
	}
}
