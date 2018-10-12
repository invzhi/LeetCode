// 90. Subsets II

// Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).
// Note: The solution set must not contain duplicate subsets.

// Example:
// Input: [1,2,2]
// Output:
// [
//   [2],
//   [1],
//   [1,2,2],
//   [2,2],
//   [1,2],
//   []
// ]

package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	sets := make([][]int, 1, 1<<uint(len(nums)))
	for i, j := 0, 0; i < len(nums); i = j {
		n := len(sets)
		for k := 0; j < len(nums) && nums[i] == nums[j]; j, k = j+1, k+n {
			for _, set := range sets[k:] {
				s := make([]int, len(set), len(set)+1)
				copy(s, set)
				sets = append(sets, append(s, nums[i]))
			}
		}
	}
	return sets
}
