// 152. Maximum Product Subarray

// Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

// Example 1:
// Input: [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.

// Example 2:
// Input: [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

package leetcode

func maxProduct(nums []int) int {
	max := nums[0]

	p1, p2 := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < 0 {
			p1, p2 = p2, p1
		}

		p1 *= num
		p2 *= num
		if num < p1 {
			p1 = num
		}
		if num > p2 {
			p2 = num
		}

		if p2 > max {
			max = p2
		}
	}

	return max
}
