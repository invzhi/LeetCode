// 450. Delete Node in a BST

// Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
// Basically, the deletion can be divided into two stages:
// 1. Search for a node to remove.
// 2. If the node is found, delete the node.
// Note: Time complexity should be O(height of tree).

// Example:
// root = [5,3,6,2,4,null,7]
// key = 3
//     5
//    / \
//   3   6
//  / \   \
// 2   4   7
// Given key to delete is 3. So we find the node with value 3 and delete it.
// One valid answer is [5,4,6,2,null,null,7], shown in the following BST.
//     5
//    / \
//   4   6
//  /     \
// 2       7
// Another valid answer is [5,2,6,null,4,null,7].
//     5
//    / \
//   2   6
//    \   \
//     4   7

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	switch {
	case key < root.Val:
		root.Left = deleteNode(root.Left, key)
	case key > root.Val:
		root.Right = deleteNode(root.Right, key)
	default:
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		root.Val = findMinNodeKey(root.Right)
		root.Right = deleteNode(root.Right, root.Val)
	}
	return root
}

func findMinNodeKey(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}
