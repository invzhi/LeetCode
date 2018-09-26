package leetcode

import (
	"strconv"
	"strings"
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
	if l == nil {
		return "<nil>"
	}

	var b strings.Builder

	b.WriteByte('(')
	b.WriteString(strconv.Itoa(l.Val))
	for l = l.Next; l != nil; l = l.Next {
		b.WriteString(" -> ")
		b.WriteString(strconv.Itoa(l.Val))
	}
	b.WriteByte(')')

	return b.String()
}
