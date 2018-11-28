// 145. Binary Tree Postorder Traversal

// Given a binary tree, return the postorder traversal of its nodes' values.

// Example:
// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
// Output: [3,2,1]

// Follow up: Recursive solution is trivial, could you do it iteratively?

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var vals []int

	var stack []*TreeNode
	for p := root; p != nil || len(stack) > 0; p = p.Right {
		if p == nil {
			n := len(stack) - 1
			p = stack[n]
			stack = stack[:n]
		}
		vals = append([]int{p.Val}, vals...)
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}

	return vals
}
