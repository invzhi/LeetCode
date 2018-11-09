// 437. Path Sum III

// You are given a binary tree in which each node contains an integer value.
// Find the number of paths that sum to a given value.
// The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).
// The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

// Example:
// root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//       10
//      /  \
//     5   -3
//    / \    \
//   3   2   11
//  / \   \
// 3  -2   1
// Return 3. The paths that sum to 8 are:
// 1.  5 -> 3
// 2.  5 -> 2 -> 1
// 3. -3 -> 11

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	return numPath(root, sum, true)
}

func numPath(root *TreeNode, sum int, top bool) int {
	if root == nil {
		return 0
	}

	var num int
	if root.Val == sum {
		num += 1
	}
	if top {
		num += numPath(root.Left, sum, true) + numPath(root.Right, sum, true)
	}
	sum -= root.Val
	num += numPath(root.Left, sum, false) + numPath(root.Right, sum, false)
	return num
}
