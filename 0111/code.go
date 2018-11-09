// 111. Minimum Depth of Binary Tree

// Given a binary tree, find its minimum depth.
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
// Note: A leaf is a node with no children.

// Example: Given binary tree [3,9,20,null,null,15,7], 3
//    / \
//   9  20
//     /  \
//    15   7
// return its depth = 2.

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	ld := minDepth(root.Left)
	rd := minDepth(root.Right)

	d := ld
	if rd < ld {
		d = rd
	}
	return d + 1
}
