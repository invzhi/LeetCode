// 687. Longest Univalue Path

// Given a binary tree, find the length of the longest path where each node in the path has the same value. This path may or may not pass through the root.
// Note: The length of path between two nodes is represented by the number of edges between them.

// Example 1:
// Input:
//     5
//    / \
//   4   5
//  / \   \
// 1   1   5
// Output:
// 2

// Example 2:
// Input:
//     1
//    / \
//   4   5
//  / \   \
// 4   4   5
// Output:
// 2

// Note: The given binary tree has not more than 10000 nodes. The height of the tree is not more than 1000.

package leetcode

// TreeNode is node of binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var length int

	var f func(*TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		a, b := f(root.Left), f(root.Right)

		var l, r int
		if root.Left != nil && root.Left.Val == root.Val {
			l = a + 1
		}
		if root.Right != nil && root.Right.Val == root.Val {
			r = b + 1
		}

		if l+r > length {
			length = l + r
		}

		if r > l {
			l = r
		}
		return l
	}

	f(root)
	return length
}
