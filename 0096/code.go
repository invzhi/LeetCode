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
	num := 1
	for i := 0; i < n; i++ {
		num = num * (4*i + 2) / (i + 2)
	}
	return num
}
