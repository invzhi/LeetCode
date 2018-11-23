// 102. Binary Tree Level Order Traversal

// Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its level order traversal as:
// [
//   [3],
//   [9,20],
//   [15,7]
// ]

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	return dfs(nil, 0, root)
}

func dfs(order [][]int, level int, root *TreeNode) [][]int {
	if root == nil {
		return order
	}

	if level == len(order) {
		order = append(order, nil)
	}
	order[level] = append(order[level], root.Val)

	order = dfs(order, level+1, root.Left)
	order = dfs(order, level+1, root.Right)

	return order
}
