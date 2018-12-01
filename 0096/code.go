// 96. Unique Binary Search Trees

// Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

// Example:
// Input: 3
// Output: 5
// Explanation:
// Given n = 3, there are a total of 5 unique BST's:
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

package leetcode

func numTrees(n int) int {
	nums := make([]int, n+1)
	nums[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			nums[i] += nums[j] * nums[i-1-j]
		}
	}
	return nums[n]
}
