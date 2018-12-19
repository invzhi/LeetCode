// 701. Insert into a Binary Search Tree

// Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.
// Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

// For example,
// Given the tree:
//         4
//        / \
//       2   7
//      / \
//     1   3
// And the value to insert: 5
// You can return this binary search tree:
//          4
//        /   \
//       2     7
//      / \   /
//     1   3 5
// This tree is also valid:
//          5
//        /   \
//       2     7
//      / \
//     1   3
//          \
//           4

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	inserted := &TreeNode{Val: val}
	if root == nil {
		return inserted
	}

	for p := root; ; {
		if val < p.Val {
			if p.Left == nil {
				p.Left = inserted
				break
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = inserted
				break
			}
			p = p.Right
		}
	}
	return root
}
