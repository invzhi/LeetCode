// 114. Flatten Binary Tree to Linked List

// Given a binary tree, flatten it to a linked list in-place.

// For example, given the following tree:
//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
// The flattened tree should look like:
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	for ; root != nil; root = root.Right {
		right := root.Right
		root.Left, root.Right = nil, root.Left

		p := root
		for p.Right != nil {
			p = p.Right
		}
		p.Right = right
	}
}
