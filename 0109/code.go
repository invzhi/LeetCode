// 109. Convert Sorted List to Binary Search Tree

// Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

// Example:
// Given the sorted linked list: [-10,-3,0,5,9],
// One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
//       0
//      / \
//    -3   9
//    /   /
//  -10  5

package leetcode

// ListNode is a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	lo, hi := head, slow.Next
	prev.Next = nil
	slow.Next = nil

	return &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST(lo),
		Right: sortedListToBST(hi),
	}
}
