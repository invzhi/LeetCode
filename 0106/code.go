// 106. Construct Binary Tree from Inorder and Postorder Traversal

// Given inorder and postorder traversal of a tree, construct the binary tree.
// Note: You may assume that duplicates do not exist in the tree.

// For example, given
// inorder = [9,3,15,20,7]
// postorder = [9,15,7,20,3]
// Return the following binary tree:
//     3
//    / \
//   9  20
//     /  \
//    15   7

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	n := len(postorder) - 1
	i := findIndex(inorder, postorder[n])
	return &TreeNode{
		Val:   postorder[n],
		Left:  buildTree(inorder[:i], postorder[:i]),
		Right: buildTree(inorder[i+1:], postorder[i:n]),
	}
}

func findIndex(order []int, v int) int {
	var i int
	for order[i] != v {
		i++
	}
	return i
}
