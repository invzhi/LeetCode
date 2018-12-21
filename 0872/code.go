// 872. Leaf-Similar Trees

// Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence.
// For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
// Two binary trees are considered leaf-similar if their leaf value sequence is the same.
// Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

// Note: Both of the given trees will have between 1 and 100 nodes.

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1, root2 *TreeNode) bool {
	var leaves1, leaves2 []int

	var f func([]int, *TreeNode) []int
	f = func(leaves []int, root *TreeNode) []int {
		if root == nil {
			return leaves
		}
		if root.Left == nil && root.Right == nil {
			return append(leaves, root.Val)
		}

		leaves = f(leaves, root.Left)
		leaves = f(leaves, root.Right)
		return leaves
	}
	leaves1 = f(leaves1, root1)
	leaves2 = f(leaves2, root2)

	if len(leaves1) != len(leaves2) {
		return false
	}
	for i := range leaves1 {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}
