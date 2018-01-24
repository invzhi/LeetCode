// 2. Add Two Numbers

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	sum := new(ListNode)
	n := sum

	var extra int

	for l1 != nil || l2 != nil {
		n.Next = new(ListNode)
		n = n.Next

		if l1 != nil {
			extra += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			extra += l2.Val
			l2 = l2.Next
		}

		n.Val = extra % 10
		extra /= 10
	}

	if extra > 0 {
		n.Next = &ListNode{Val: extra}
	}

	return sum.Next
}
