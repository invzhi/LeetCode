// 222. Count Complete Tree Nodes

// Given a complete binary tree, count the number of nodes.
// Note:
// Definition of a complete binary tree from Wikipedia:
// In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2^h nodes inclusive at the last level h.

// Example:
// Input:
//     1
//    / \
//   2   3
//  / \  /
// 4  5 6
// Output: 6

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var cnt int
	for h := height(root) - 1; root != nil; h-- {
		rh := height(root.Right)

		if h == rh {
			cnt += 1 << uint(h)
			root = root.Right
			h = rh
		} else {
			cnt += 1 << uint(rh)
			root = root.Left
		}
	}
	return cnt
}

func height(root *TreeNode) int {
	var h int
	for p := root; p != nil; p = p.Left {
		h++
	}
	return h
}
