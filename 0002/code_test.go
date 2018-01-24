package leetcode

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func asList(vs ...int) *ListNode {
	l := new(ListNode)
	n := l
	for _, v := range vs {
		n.Next = &ListNode{Val: v}
		n = n.Next
	}
	return l.Next
}

func (l *ListNode) String() string {
	buf := bytes.NewBufferString("(")
	if l != nil {
		buf.WriteString(strconv.Itoa(l.Val))
		l = l.Next
	}
	for l != nil {
		buf.WriteString(" -> ")
		buf.WriteString(strconv.Itoa(l.Val))
		l = l.Next
	}
	buf.WriteString(")")
	return buf.String()
}

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		l1, l2 *ListNode
		sum    *ListNode
	}{
		{asList(2, 4, 6), asList(5, 6, 4), asList(7, 0, 1, 1)},
		{asList(2, 4, 3), asList(5, 6, 4), asList(7, 0, 8)},
		{asList(2, 4), asList(5, 6, 4), asList(7, 0, 5)},
		{asList(2), asList(5, 6, 4), asList(7, 6, 4)},
		{asList(2), asList(5, 6), asList(7, 6)},
		{asList(2), asList(5), asList(7)},
		{asList(), asList(5), asList(5)},
		{asList(), asList(), asList()},
	}

	for _, tt := range tests {
		sum := addTwoNumbers(tt.l1, tt.l2)
		if reflect.DeepEqual(sum, tt.sum) == false {
			t.Errorf("addTwoNumbers(%s, %s) return %s, want %s", tt.l1, tt.l2, sum, tt.sum)
		}
	}
}
