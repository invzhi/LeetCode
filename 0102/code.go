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

// TreeNode is node of binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	order := dfs(root, nil, 0)
	return order
}

func dfs(root *TreeNode, order [][]int, level int) [][]int {
	if len(order) <= level {
		order = append(order, nil)
	}
	order[level] = append(order[level], root.Val)

	if root.Left != nil {
		order = dfs(root.Left, order, level+1)
	}

	if root.Right != nil {
		order = dfs(root.Right, order, level+1)
	}

	return order
}