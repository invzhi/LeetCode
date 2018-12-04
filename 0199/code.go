// 199. Binary Tree Right Side View

// Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

// Example:
// Input: [1,2,3,null,5,null,4]
// Output: [1, 3, 4]
// Explanation:
//    1            <---
//  /   \
// 2     3         <---
//  \     \
//   5     4       <---

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var values []int

	var f func(*TreeNode, int) int
	f = func(root *TreeNode, depth int) int {
		if root == nil {
			return depth
		}

		if depth == 0 {
			values = append(values, root.Val)
			depth = f(root.Right, 0)
		} else {
			depth = f(root.Right, depth-1)
		}
		return f(root.Left, depth) + 1
	}
	f(root, 0)

	return values
}
