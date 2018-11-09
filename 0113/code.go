// 113. Path Sum II

// Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
// Note: A leaf is a node with no children.

// Example:
// Given the below binary tree and sum = 22,
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \    / \
// 7    2  5   1
// Return:
// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val != sum {
			return nil
		}
		return [][]int{{root.Val}}
	}

	var paths [][]int
	sum -= root.Val
	for _, path := range pathSum(root.Left, sum) {
		path = append([]int{root.Val}, path...)
		paths = append(paths, path)
	}
	for _, path := range pathSum(root.Right, sum) {
		path = append([]int{root.Val}, path...)
		paths = append(paths, path)
	}
	return paths
}
