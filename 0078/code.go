// 78. Subsets

// Given a set of distinct integers, nums, return all possible subsets (the power set).
// Note: The solution set must not contain duplicate subsets.

// Example:
// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

package leetcode

func subsets(nums []int) [][]int {
	sets := make([][]int, 0, 1<<uint(len(nums)))
	set := make([]int, 0, len(nums))
	return dfs(sets, set, nums)
}

func dfs(sets [][]int, set, nums []int) [][]int {
	if len(nums) == 0 {
		s := make([]int, len(set))
		copy(s, set)
		return append(sets, s)
	}

	sets = dfs(sets, set, nums[1:])
	set = append(set, nums[0])
	return dfs(sets, set, nums[1:])
}
