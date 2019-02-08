// 147. Insertion Sort List

// Sort a linked list using insertion sort.

// Algorithm of Insertion Sort:
// 1. Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
// 2. At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
// 3. It repeats until no input elements remain.

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	l := &ListNode{Next: head}
	for prev, p := head, head.Next; p != nil; p = p.Next {
		if prev.Val <= p.Val {
			prev = p
			continue
		}

		q := l
		for q.Next.Val < p.Val {
			q = q.Next
		}
		prev.Next = p.Next
		p.Next = q.Next
		q.Next = p
		p = prev
	}
	return l.Next
}
