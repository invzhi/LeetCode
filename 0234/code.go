// 234. Palindrome Linked List

// Given a singly linked list, determine if it is a palindrome.

// Example 1:
// Input: 1->2
// Output: false

// Example 2:
// Input: 1->2->2->1
// Output: true

// Follow up:
// Could you do it in O(n) time and O(1) space?

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev, slow = slow, next
	}

	prev, head = slow, prev
	if fast != nil {
		slow = slow.Next
	}

	palindrome := true
	for head != nil {
		if head.Val != slow.Val {
			palindrome = false
		}
		next := head.Next
		head.Next = prev
		prev, head = head, next

		slow = slow.Next
	}
	return palindrome
}
