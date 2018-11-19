// 61. Rotate List

// Given a linked list, rotate the list to the right by k places, where k is non-negative.

// Example 1:
// Input: 1->2->3->4->5->NULL, k = 2
// Output: 4->5->1->2->3->NULL
// Explanation:
// rotate 1 steps to the right: 5->1->2->3->4->NULL
// rotate 2 steps to the right: 4->5->1->2->3->NULL

// Example 2:
// Input: 0->1->2->NULL, k = 4
// Output: 2->0->1->NULL
// Explanation:
// rotate 1 steps to the right: 2->0->1->NULL
// rotate 2 steps to the right: 1->2->0->NULL
// rotate 3 steps to the right: 0->1->2->NULL
// rotate 4 steps to the right: 2->0->1->NULL

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var n int
	var p, q *ListNode
	for q := head; q != nil; q = q.Next {
		p = q
		n++
	}
	p.Next, q = head, head

	bias := n - k%n
	for i := 0; i < bias; i++ {
		p, q = q, q.Next
	}
	p.Next = nil
	return q
}
