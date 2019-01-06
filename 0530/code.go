// 530. Minimum Absolute Difference in BST

// Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.

// Example:
// Input:
//    1
//     \
//      3
//     /
//    2
// Output:
// 1

// Explanation: The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).

// Note: There are at least two nodes in this BST.

package leetcode

import "math"

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt64

	var prev *TreeNode
	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}

		f(root.Left)
		if prev != nil && root.Val-prev.Val < min {
			min = root.Val - prev.Val
		}
		prev = root
		f(root.Right)
	}
	f(root)

	return min
}
