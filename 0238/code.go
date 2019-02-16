// 238. Product of Array Except Self

// Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

// Example:
// Input:  [1,2,3,4]
// Output: [24,12,8,6]

// Note: Please solve it without division and in O(n).

// Follow up: Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)

package leetcode

func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))

	product := 1
	for i := range products {
		products[i] = product
		product *= nums[i]
	}

	product = 1
	for i := len(products) - 1; i >= 0; i-- {
		products[i] *= product
		product *= nums[i]
	}

	return products
}
