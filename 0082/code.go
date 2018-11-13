// 82. Remove Duplicates from Sorted List II

// Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

// Example 1:
// Input: 1->2->3->3->4->4->5
// Output: 1->2->5

// Example 2:
// Input: 1->1->1->2->3
// Output: 2->3

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	l := &ListNode{Next: head}
	for prev := l; head != nil; {
		if head.Next != nil && head.Val == head.Next.Val {
			for val := head.Val; head != nil && head.Val == val; {
				dup := head
				head = head.Next
				dup.Next = nil
			}
			prev.Next = head
		} else {
			prev = head
			head = head.Next
		}
	}
	return l.Next
}
