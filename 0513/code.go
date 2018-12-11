// 513. Find Bottom Left Tree Value

// Given a binary tree, find the leftmost value in the last row of the tree.

// Example 1:
// Input:
//     2
//    / \
//   1   3
// Output:
// 1

// Example 2:
// Input:
//         1
//        / \
//       2   3
//      /   / \
//     4   5   6
//        /
//       7
// Output:
// 7

// Note: You may assume the tree (i.e., the given root node) is not NULL.

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var value, depth int

	var f func(*TreeNode, int)
	f = func(root *TreeNode, d int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && d > depth {
			depth = d
			value = root.Val
		}

		f(root.Left, d+1)
		f(root.Right, d+1)
	}
	f(root, 1)

	return value
}
