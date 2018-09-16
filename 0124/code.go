// 124. Binary Tree Maximum Path Sum

// Given a non-empty binary tree, find the maximum path sum.
// For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

// Example 1:
// Input: [1,2,3]
//     1
//    / \
//   2   3
// Output: 6

// Example 2:
// Input: [-10,9,20,null,null,15,7]
//    -10
//    / \
//   9  20
//     /  \
//    15   7
// Output: 42

package leetcode

// TreeNode is node of binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val

	var f func(*TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		sum := root.Val

		l, r := f(root.Left), f(root.Right)
		if l > 0 {
			sum += l
		}
		if r > 0 {
			sum += r
		}

		if sum > maxSum {
			maxSum = sum
		}

		var max int
		if l > max {
			max = l
		}
		if r > max {
			max = r
		}
		return root.Val + max
	}

	f(root)
	return maxSum
}
