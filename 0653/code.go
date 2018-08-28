// 653. Two Sum IV - Input is a BST

// Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

// Example 1:
// Input:
//     5
//    / \
//   3   6
//  / \   \
// 2   4   7
// Target = 9
// Output: True

// Example 2:
// Input:
//     5
//    / \
//   3   6
//  / \   \
// 2   4   7
// Target = 28
// Output: False

package leetcode

// TreeNode is node of binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return search(root, k, m)
}

func search(node *TreeNode, k int, m map[int]struct{}) bool {
	if node == nil {
		return false
	}
	if _, ok := m[k-node.Val]; ok {
		return true
	}
	m[node.Val] = struct{}{}
	return search(node.Left, k, m) || search(node.Right, k, m)
}
