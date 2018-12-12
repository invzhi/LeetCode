// 515. Find Largest Value in Each Tree Row

// You need to find the largest value in each row of a binary tree.

// Example:
// Input:
//           1
//          / \
//         3   2
//        / \   \
//       5   3   9
// Output: [1, 3, 9]

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var vals []int

	var f func(*TreeNode, int)
	f = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if depth == len(vals) {
			vals = append(vals, root.Val)
		} else if root.Val > vals[depth] {
			vals[depth] = root.Val
		}

		f(root.Left, depth+1)
		f(root.Right, depth+1)
	}
	f(root, 0)

	return vals
}
