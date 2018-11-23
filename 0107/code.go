// 107. Binary Tree Level Order Traversal II

// Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its bottom-up level order traversal as:
// [
//   [15,7],
//   [9,20],
//   [3]
// ]

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	return dfs(nil, 0, root)
}

func dfs(order [][]int, level int, root *TreeNode) [][]int {
	if root == nil {
		return order
	}

	if level == len(order) {
		order = append([][]int{nil}, order...)
	}
	i := len(order) - 1 - level
	order[i] = append(order[i], root.Val)

	order = dfs(order, level+1, root.Left)
	order = dfs(order, level+1, root.Right)

	return order
}
