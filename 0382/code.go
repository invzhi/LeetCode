// 382. Linked List Random Node

// Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.

// Follow up:
// What if the linked list is extremely large and its length is unknown to you? Could you solve this efficiently without using extra space?

// Example:
// ListNode head = new ListNode(1);
// head.next = new ListNode(2);
// head.next.next = new ListNode(3);
// Solution solution = new Solution(head);
// solution.getRandom();

package leetcode

import "math/rand"

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution is the solution for the reservoir sampling algorithm.
type Solution struct {
	head *ListNode
}

// Constructor accepts a head node of a linked list.
// Note that the head is guaranteed to be not null, so it contains at least one node.
func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

// GetRandom returns a random node's value.
func (s *Solution) GetRandom() int {
	val := s.head.Val

	i := 1
	for p := s.head.Next; p != nil; p = p.Next {
		if rand.Intn(i+1) == 0 {
			val = p.Val
		}
		i++
	}
	return val
}
