// 283. Move Zeroes

// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Example:
// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Note:
// 1. You must do this in-place without making a copy of the array.
// 2. Minimize the total number of operations.

package leetcode

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
