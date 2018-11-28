// 144. Binary Tree Preorder Traversal

// Given a binary tree, return the preorder traversal of its nodes' values.

// Example:
// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
// Output: [1,2,3]

// Follow up: Recursive solution is trivial, could you do it iteratively?

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var vals []int

	var stack []*TreeNode
	for p := root; p != nil || len(stack) > 0; p = p.Left {
		if p == nil {
			n := len(stack) - 1
			p = stack[n]
			stack = stack[:n]
		}
		vals = append(vals, p.Val)
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
	}

	return vals
}
