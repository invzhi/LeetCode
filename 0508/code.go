// 508. Most Frequent Subtree Sum

// Given the root of a tree, you are asked to find the most frequent subtree sum. The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie, return all the values with the highest frequency in any order.

// Examples 1
// Input:
//   5
//  / \
// 2  -3
// return [2, -3, 4], since all the values happen only once, return all of them in any order.

// Examples 2
// Input:
//   5
//  / \
// 2  -5
// return [2], since 2 happens twice, however -5 only occur once.

// Note: You may assume the sum of values in any subtree is in the range of 32-bit signed integer.

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	var max int
	cnt := make(map[int]int)

	var f func(*TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		sum := root.Val
		sum += f(root.Left)
		sum += f(root.Right)

		cnt[sum]++
		if cnt[sum] > max {
			max = cnt[sum]
		}

		return sum
	}
	f(root)

	var sums []int
	for sum, c := range cnt {
		if c == max {
			sums = append(sums, sum)
		}
	}

	return sums
}
