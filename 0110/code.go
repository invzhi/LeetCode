// 110. Balanced Binary Tree

// Given a binary tree, determine if it is height-balanced.
// For this problem, a height-balanced binary tree is defined as: a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

// Example 1:
// Given the following tree [3,9,20,null,null,15,7]:
//     3
//    / \
//   9  20
//     /  \
//    15   7
// Return true.

// Example 2:
// Given the following tree [1,2,2,3,3,null,null,4,4]:
//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4
// Return false.

package leetcode

// TreeNode is node of binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := height(root)
	return balanced
}

func height(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	left, ok := height(root.Left)
	if !ok {
		return 0, false
	}
	right, ok := height(root.Right)
	if !ok {
		return 0, false
	}

	if sub := left - right; sub < -1 || sub > 1 {
		return 0, false
	}

	h := left
	if right > left {
		h = right
	}
	return h + 1, true
}
