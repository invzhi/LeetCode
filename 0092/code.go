// 92. Reverse Linked List II

// Reverse a linked list from position m to n. Do it in one-pass.

// Note: 1 ≤ m ≤ n ≤ length of list.

// Example:
// Input: 1->2->3->4->5->NULL, m = 2, n = 4
// Output: 1->4->3->2->5->NULL

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m, n int) *ListNode {
	l := &ListNode{Next: head}
	prev := l
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	head = prev.Next

	p, q := prev.Next, prev.Next.Next
	for i := m; i < n; i++ {
		next := q.Next
		q.Next = p
		p, q = q, next
	}
	prev.Next, head.Next = p, q
	return l.Next
}
