package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

func (l MyLinkedList) String() string {
	node := l.head
	if node == nil {
		return "<nil>"
	}

	var b strings.Builder

	b.WriteByte('(')
	b.WriteString(strconv.Itoa(node.val))
	for node = node.next; node != nil; node = node.next {
		b.WriteString(" -> ")
		b.WriteString(strconv.Itoa(node.val))
	}
	b.WriteByte(')')

	return b.String()
}

func TestLinkedList(t *testing.T) {
	l := Constructor()

	vs := []int{1, 2, 3, -1}

	// add a head node, then delete it.
	l.AddAtHead(vs[0])
	t.Logf("l.AddAtHead(%v) return %s", vs[0], l)
	if v := l.Get(0); v != vs[0] {
		t.Errorf("l.Get(%v) return %v, want %v", 0, v, vs[0])
	}

	l.DeleteAtIndex(0)
	t.Logf("l.DeleteAtIndex(0) return %s", l)
	if l.head != nil || l.tail != nil || l.length != 0 {
		t.Errorf("l is not a empty linked list when its all nodes deleted")
	}

	// add two tail node, then insert a middle node.
	l.AddAtTail(vs[0])
	t.Logf("l.AddAtTail(%v) return %s", vs[0], l)
	if v := l.Get(l.length - 1); v != vs[0] {
		t.Errorf("l.Get(%v) return %v, want %v", l.length-1, v, vs[0])
	}

	l.AddAtTail(vs[2])
	t.Logf("l.AddAtTail(%v) return %s", vs[2], l)
	if v := l.Get(l.length - 1); v != vs[2] {
		t.Errorf("l.Get(%v) return %v, want %v", l.length-1, v, vs[2])
	}

	l.AddAtIndex(-1, 2)
	t.Logf("l.AddAtIndex(%v, %v) return %s", -1, 2, l)
	if v := l.Get(l.length - 1); v != vs[2] {
		t.Errorf("l.Get(%v) return %v, want %v", l.length-1, v, vs[2])
	}

	l.AddAtIndex(1, vs[1])
	t.Logf("l.AddAtIndex(%v, %v) return %s", 1, vs[1], l)
	if v := l.Get(1); v != vs[1] {
		t.Errorf("l.Get(%v) return %v, want %v", 1, v, vs[1])
	}

	for i, v := range vs {
		if v != l.Get(i) {
			t.Errorf("l.Get(%v) return %v, want %v", i, l.Get(i), v)
		}
	}

	l.DeleteAtIndex(-1)
	t.Logf("l.DeleteAtIndex(%v) return %s", -1, l)
	l.DeleteAtIndex(0)
	t.Logf("l.DeleteAtIndex(%v) return %s", 0, l)
	l.DeleteAtIndex(1)
	t.Logf("l.DeleteAtIndex(%v) return %s", 1, l)
	l.DeleteAtIndex(0)
	t.Logf("l.DeleteAtIndex(%v) return %s", 0, l)

	for i, v := range vs[:3] {
		l.AddAtIndex(i, v)
		t.Logf("l.AddAtIndex(%v, %v) return %s", i, v, l)
		if v != l.Get(i) {
			t.Errorf("l.Get(%v) return %v, want %v", i, l.Get(i), v)
		}
	}
}
