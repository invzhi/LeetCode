// 652. Find Duplicate Subtrees

// Given a binary tree, return all duplicate subtrees. For each kind of duplicate subtrees, you only need to return the root node of any one of them.
// Two trees are duplicate if they have the same structure with same node values.

// Example 1:
//         1
//        / \
//       2   3
//      /   / \
//     4   2   4
//        /
//       4
// The following are two duplicate subtrees:
//       2
//      /
//     4
// and
//     4

package leetcode

import "strconv"

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var subtrees []*TreeNode
	cnt := make(map[string]int)

	var f func(*TreeNode) string
	f = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}

		str := strconv.Itoa(root.Val) + "," + f(root.Left) + "," + f(root.Right)

		cnt[str]++
		if cnt[str] == 2 {
			subtrees = append(subtrees, root)
		}
		return str
	}
	f(root)

	return subtrees
}
