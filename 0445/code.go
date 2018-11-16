// 445. Add Two Numbers II

// You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Follow up:
// What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

// Example:
// Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 8 -> 0 -> 7

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	n1, n2 := count(l1), count(l2)
	if n1 < n2 {
		n1, n2 = n2, n1
		l1, l2 = l2, l1
	}

	l := new(ListNode)
	prev, n := l, l
	for i := 0; i < n1-n2; i++ {
		n.Next = &ListNode{Val: l1.Val}
		n = n.Next
		l1 = l1.Next
		if n.Val < 9 {
			prev = n
		}
	}

	for l1 != nil {
		n.Next = &ListNode{Val: l1.Val + l2.Val}
		n = n.Next
		if n.Val < 9 {
			prev = n
		} else if n.Val > 9 {
			prev.Val++
			for prev = prev.Next; prev != n; prev = prev.Next {
				prev.Val = 0
			}
			n.Val -= 10
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l.Val == 0 {
		l = l.Next
	}
	return l
}

func count(l *ListNode) int {
	var cnt int
	for l != nil {
		l = l.Next
		cnt++
	}
	return cnt
}
