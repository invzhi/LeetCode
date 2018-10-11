// 77. Combinations

// Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

// Example:
// Input: n = 4, k = 2
// Output:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]

package leetcode

func combine(n int, k int) [][]int {
	var combinations [][]int

	var f func([]int, int)
	f = func(nums []int, index int) {
		if len(nums) == k {
			n := make([]int, k)
			copy(n, nums)
			combinations = append(combinations, n)
			return
		}

		for i := index; i <= n; i++ {
			f(append(nums, i), i+1)
		}
	}

	f(make([]int, 0, k), 1)
	return combinations
}
