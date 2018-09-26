// 39. Combination Sum

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
// The same repeated number may be chosen from candidates unlimited number of times.
// Note:
// - All numbers (including target) will be positive integers.
// - The solution set must not contain duplicate combinations.

// Example 1:
// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]

// Example 2:
// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var combinations [][]int

	var f func([]int, int, int)
	f = func(combination []int, index int, target int) {
		if target == 0 {
			n := make([]int, len(combination))
			copy(n, combination)
			combinations = append(combinations, n)
			return
		}

		for i, c := range candidates[index:] {
			if c > target {
				continue
			}

			combination = append(combination, c)
			f(combination, index+i, target-c)
			combination = combination[:len(combination)-1]
		}
	}

	f(nil, 0, target)
	return combinations
}
