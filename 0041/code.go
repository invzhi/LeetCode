// 41. First Missing Positive

// Given an unsorted integer array, find the smallest missing positive integer.

// Example 1:
// Input: [1,2,0]
// Output: 3

// Example 2:
// Input: [3,4,-1,1]
// Output: 2

// Example 3:
// Input: [7,8,9,11,12]
// Output: 1

// Note:
// Your algorithm should run in O(n) time and uses constant extra space.

package leetcode

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for j := nums[i] - 1; j >= 0 && j < len(nums); j = nums[i] - 1 {
			if i == j || nums[i] == nums[j] {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	var i int
	for i < len(nums) && nums[i] == i+1 {
		i++
	}
	return i + 1
}
