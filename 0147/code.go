// 147. Insertion Sort List

// Sort a linked list using insertion sort.

// Algorithm of Insertion Sort:
// 1. Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
// 2. At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
// 3. It repeats until no input elements remain.

package leetcode

// ListNode is node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	l := &ListNode{Next: head}
	for prev, i := head, head.Next; i != nil; i = i.Next {
		if prev.Val <= i.Val {
			prev = i
			continue
		}

		j := l
		for j.Next.Val <= i.Val {
			j = j.Next
		}
		prev.Next = i.Next
		i.Next = j.Next
		j.Next = i
		i = prev
	}
	return l.Next
}
