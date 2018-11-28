// 94. Binary Tree Inorder Traversal

// Given a binary tree, return the inorder traversal of its nodes' values.

// Example:
// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
// Output: [1,3,2]

// Follow up: Recursive solution is trivial, could you do it iteratively?

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var vals []int

	var stack []*TreeNode
	for p := root; p != nil || len(stack) > 0; p = p.Right {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		n := len(stack) - 1
		p = stack[n]
		stack = stack[:n]

		vals = append(vals, p.Val)
	}

	return vals
}
