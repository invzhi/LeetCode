// 24. Swap Nodes in Pairs

// Given a linked list, swap every two adjacent nodes and return its head.

// Example:
// Given 1->2->3->4, you should return the list as 2->1->4->3.

// Note:
// - Your algorithm should use only constant extra space.
// - You may not modify the values in the list's nodes, only nodes itself may be changed.

package leetcode

// ListNode is node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	for head != nil && head.Next != nil {
		first, second := head, head.Next

		next := second.Next
		if next != nil && next.Next != nil {
			next = next.Next
		}

		head = second.Next
		first.Next, second.Next = next, first
	}
	return newHead
}
