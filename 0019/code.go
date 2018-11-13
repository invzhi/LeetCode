// 19. Remove Nth Node From End of List

// Given a linked list, remove the n-th node from the end of list and return its head.

// Example:
// Given linked list: 1->2->3->4->5, and n = 2.
// After removing the second node from the end, the linked list becomes 1->2->3->5.

// Note:
// Given n will always be valid.

// Follow up:
// Could you do this in one pass?

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := &ListNode{Next: head}
	for i := 0; i < n; i++ {
		head = head.Next
	}

	prev := l
	for head != nil {
		prev = prev.Next
		head = head.Next
	}
	prev.Next = prev.Next.Next
	return l.Next
}
