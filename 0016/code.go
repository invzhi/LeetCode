// 16. 3Sum Closest

// Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

// Example: Given array nums = [-1, 2, 1, -4], and target = 1.

// The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			switch {
			case sum < target:
				lo++
			case sum > target:
				hi--
			default:
				return sum
			}

			if diff(sum, target) < diff(closest, target) {
				closest = sum
			}
		}
	}
	return closest
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
