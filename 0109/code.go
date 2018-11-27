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
	var length int
	for p := head; p != nil; p = p.Next {
		length++
	}

	var f func(int) *TreeNode
	f = func(n int) *TreeNode {
		if n == 0 {
			return nil
		}
		n--
		mid := n / 2

		node := new(TreeNode)

		node.Left = f(n - mid)
		node.Val = head.Val
		head = head.Next
		node.Right = f(mid)

		return node
	}
	return f(length)
}
