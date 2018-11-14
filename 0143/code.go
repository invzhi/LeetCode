// 143. Reorder List

// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→L(n-1)→L2→L(n-2)→…

// You may not modify the values in the list's nodes, only nodes itself may be changed.

// Example 1:
// Given 1->2->3->4, reorder it to 1->4->2->3.

// Example 2:
// Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev, slow = slow, next
	}

	for p, q := head, prev; p != q; p, q = q, p {
		next := p.Next
		p.Next = q
		p = next
	}
}
