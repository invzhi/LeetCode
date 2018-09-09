// 33. Search in Rotated Sorted Array

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// You are given a target value to search. If found in the array return its index, otherwise return -1.
// You may assume no duplicate exists in the array.
// Your algorithm's runtime complexity must be in the order of O(log n).

// Example 1:
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4

// Example 2:
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1

package leetcode

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	zero := findZeroIndex(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2

		i := (zero + mid) % n
		switch {
		case nums[i] < target:
			lo = mid + 1
		case nums[i] > target:
			hi = mid - 1
		default:
			return i
		}
	}
	return -1
}

func findZeroIndex(nums []int) int {
	var zero int
	lo, hi := 0, len(nums)-1
	if nums[lo] <= nums[hi] {
		return 0
	}
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[mid+1] {
			zero = mid + 1
			break
		} else if nums[mid] >= nums[lo] {
			lo = mid + 1
		} else if nums[mid] <= nums[hi] {
			hi = mid - 1
		}
	}
	return zero
}
