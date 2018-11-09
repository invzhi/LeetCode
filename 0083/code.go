// 83. Remove Duplicates from Sorted List

// Given a sorted linked list, delete all duplicates such that each element appear only once.

// Example 1:
// Input: 1->1->2
// Output: 1->2

// Example 2:
// Input: 1->1->2->3->3
// Output: 1->2->3

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	first := head
	for first != nil && first.Next != nil {
		second := first.Next
		if first.Val == second.Val {
			first.Next = second.Next
			second.Next = nil
		} else {
			first = first.Next
		}
	}
	return head
}
