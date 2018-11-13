// 86. Partition List

// Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in each of the two partitions.

// Example:
// Input: head = 1->4->3->2->5->2, x = 3
// Output: 1->2->2->4->3->5

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	l := new(ListNode)
	less, gt := l, &ListNode{Next: head}
	for prev := gt; head != nil; head = head.Next {
		if head.Val < x {
			less.Next = head
			less = less.Next
			prev.Next = head.Next
		} else {
			prev = head
		}
	}
	less.Next = gt.Next
	return l.Next
}
