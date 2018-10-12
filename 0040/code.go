// 40. Combination Sum II

// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
// Each number in candidates may only be used once in the combination.
// Note:
// - All numbers (including target) will be positive integers.
// - The solution set must not contain duplicate combinations.

// Example 1:
// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]

// Example 2:
// Input: candidates = [2,5,2,1,2], target = 5,
// A solution set is:
// [
//   [1,2,2],
//   [5]
// ]

package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var combinations [][]int

	var f func([]int, int, int)
	f = func(combination []int, index int, target int) {
		if target == 0 {
			n := make([]int, len(combination))
			copy(n, combination)
			combinations = append(combinations, n)
			return
		}

		for i, j := index, index+1; i < len(candidates); i, j = j, j+1 {
			for j < len(candidates) && candidates[i] == candidates[j] {
				j++
			}

			c := combination
			var sum int
			for k := i; k < j; k++ {
				sum += candidates[k]
				if sum > target {
					break
				}
				c = append(c, candidates[k])
				f(c, j, target-sum)
			}
		}
	}

	f(nil, 0, target)
	return combinations
}
