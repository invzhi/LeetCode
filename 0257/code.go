// 257. Binary Tree Paths

// Given a binary tree, return all root-to-leaf paths.
// Note: A leaf is a node with no children.

// Example:
// Input:
//    1
//  /   \
// 2     3
//  \
//   5
// Output: ["1->2->5", "1->3"]
// Explanation: All root-to-leaf paths are: 1->2->5, 1->3

package leetcode

import (
	"strconv"
	"strings"
)

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	v := strconv.Itoa(root.Val)
	paths := append(binaryTreePaths(root.Left), binaryTreePaths(root.Right)...)

	var b strings.Builder
	for i, path := range paths {
		b.Reset()
		b.WriteString(v)
		b.WriteString("->")
		b.WriteString(path)

		paths[i] = b.String()
	}
	return paths
}
