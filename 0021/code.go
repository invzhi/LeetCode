// 21. Merge Two Sorted Lists

// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

// Example:
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

package leetcode

// ListNode is node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	l := new(ListNode)

	n := l
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			n.Next = l1
			n = n.Next
			l1 = l1.Next
		} else {
			n.Next = l2
			n = n.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		n.Next = l1
	}
	if l2 != nil {
		n.Next = l2
	}

	return l.Next
}
