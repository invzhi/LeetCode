// 203. Remove Linked List Elements

// Remove all elements from a linked list of integers that have value val.

// Example:
// Input:  1->2->6->3->4->5->6, val = 6
// Output: 1->2->3->4->5

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	l := &ListNode{Next: head}
	for prev := l; head != nil; head = head.Next {
		if head.Val == val {
			prev.Next = head.Next
		} else {
			prev = head
		}
	}
	return l.Next
}
