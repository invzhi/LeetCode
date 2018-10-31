// 47. Permutations II

// Given a collection of numbers that might contain duplicates, return all possible unique permutations.

// Example:
//
// Input: [1,1,2]
// Output:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return dfs(nil, nums, 0)
}

func dfs(permutations [][]int, nums []int, index int) [][]int {
	if index == len(nums)-1 {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		return append(permutations, permutation)
	}

	permutations = dfs(permutations, nums, index+1)
	for i := index + 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			nums[i], nums[index] = nums[index], nums[i]
			permutations = dfs(permutations, nums, index+1)
		}
	}
	max := nums[index]
	var i int
	for i = index; i+1 < len(nums); i++ {
		nums[i] = nums[i+1]
	}
	nums[i] = max
	return permutations
}
