// 154. Find Minimum in Rotated Sorted Array II

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
// Find the minimum element.
// The array may contain duplicates.

// Example 1:
// Input: [1,3,5]
// Output: 1

// Example 2:
// Input: [2,2,2,0,1]
// Output: 0

// Note:
// - This is a follow up problem to Find Minimum in Rotated Sorted Array.
// - Would allow duplicates affect the run-time complexity? How and why?

package leetcode

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] > nums[hi]:
			lo = mid + 1
		case nums[mid] < nums[hi]:
			hi = mid
		default:
			hi--
		}
	}
	return nums[lo]
}
