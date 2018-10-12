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
	sets := make([][]int, 1, 1<<uint(len(nums)))
	for _, num := range nums {
		for _, set := range sets {
			s := make([]int, len(set), len(set)+1)
			copy(s, set)
			sets = append(sets, append(s, num))
		}
	}
	return sets
}
