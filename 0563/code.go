// 563. Binary Tree Tilt

// Given a binary tree, return the tilt of the whole tree.
// The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.
// The tilt of the whole tree is defined as the sum of all nodes' tilt.

// Example:
// Input:
//          1
//        /   \
//       2     3
// Output: 1
// Explanation:
// Tilt of node 2 : 0
// Tilt of node 3 : 0
// Tilt of node 1 : |2-3| = 1
// Tilt of binary tree : 0 + 0 + 1 = 1

// Note:
// 1. The sum of node values in any subtree won't exceed the range of 32-bit integer.
// 2. All the tilt values won't exceed the range of 32-bit integer.

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var tilt int

	var f func(*TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := f(root.Left)
		r := f(root.Right)
		if l < r {
			tilt += r - l
		} else {
			tilt += l - r
		}
		return root.Val + l + r
	}
	f(root)

	return tilt
}
