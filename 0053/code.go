// 53. Maximum Subarray

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// Example:
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

// Follow up:
// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

package leetcode

func maxSubArray(nums []int) int {
	var max, sum int
	for i, num := range nums {
		sum += num

		if sum > max || i == 0 {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
