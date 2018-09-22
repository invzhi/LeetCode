// 153. Find Minimum in Rotated Sorted Array

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
// Find the minimum element.
// You may assume no duplicate exists in the array.
//
// Example 1:
// Input: [3,4,5,1,2]
// Output: 1
//
// Example 2:
// Input: [4,5,6,7,0,1,2]
// Output: 0

package leetcode

func findMin(nums []int) int {
	var min int
	lo, hi := 0, len(nums)-1
	if nums[lo] <= nums[hi] {
		return nums[lo]
	}
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[mid+1] {
			min = nums[mid+1]
			break
		} else if nums[mid] >= nums[lo] {
			lo = mid + 1
		} else if nums[mid] <= nums[hi] {
			hi = mid - 1
		}
	}
	return min
}
