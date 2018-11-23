// 637. Average of Levels in Binary Tree

// Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

// Example 1:
// Input:
//     3
//    / \
//   9  20
//     /  \
//    15   7
// Output: [3, 14.5, 11]
// Explanation:
// The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].

// Note:
// 1. The range of node's value is in the range of 32-bit signed integer.

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	sums, cnts := dfs(nil, nil, 0, root)
	for i, cnt := range cnts {
		sums[i] /= float64(cnt)
	}
	return sums
}

func dfs(sums []float64, cnts []int, level int, root *TreeNode) ([]float64, []int) {
	if root == nil {
		return sums, cnts
	}

	if level == len(sums) {
		sums = append(sums, 0)
		cnts = append(cnts, 0)
	}
	sums[level] += float64(root.Val)
	cnts[level]++

	sums, cnts = dfs(sums, cnts, level+1, root.Left)
	sums, cnts = dfs(sums, cnts, level+1, root.Right)

	return sums, cnts
}
