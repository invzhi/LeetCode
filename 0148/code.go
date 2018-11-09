// 148. Sort List

// Sort a linked list in O(n log n) time using constant space complexity.

// Example 1:
// Input: 4->2->1->3
// Output: 1->2->3->4

// Example 2:
// Input: -1->5->3->4->0
// Output: -1->0->3->4->5

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	m := findMiddle(head)

	r := sortList(m.Next)
	m.Next = nil
	l := sortList(head)

	return merge(l, r)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
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
