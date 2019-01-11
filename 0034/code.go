// 34. Find First and Last Position of Element in Sorted Array

// Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
// Your algorithm's runtime complexity must be in the order of O(log n).
// If the target is not found in the array, return [-1, -1].

// Example 1:
// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]

// Example 2:
// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]

package leetcode

func searchRange(nums []int, target int) []int {
	position := []int{-1, -1}

	hi := len(nums) - 1
	position[0] = pos1(nums, target, 0, hi)
	if position[0] == -1 {
		return position
	}

	position[1] = pos2(nums, target, position[0], hi)
	return position
}

func pos1(nums []int, target, lo, hi int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] < target:
			lo = mid + 1
		case nums[mid] > target:
			hi = mid - 1
		default:
			hi = mid
		}
	}
	if hi >= 0 && nums[hi] == target {
		return hi
	}
	return -1
}

func pos2(nums []int, target, lo, hi int) int {
	for lo < hi {
		mid := hi - (hi-lo)/2
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid
		}
	}
	return lo
}
