// 260. Single Number III

// Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

// Example:
// Input:  [1,2,1,3,2,5]
// Output: [3,5]

// Note:
// The order of the result is not important. So in the above example, [5, 3] is also correct.
// Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

package leetcode

func singleNumber(nums []int) []int {
	var single int
	for _, num := range nums {
		single ^= num
	}

	diff := single & (-single)
	once := make([]int, 2)
	for _, num := range nums {
		if num&diff == 0 {
			once[0] ^= num
		} else {
			once[1] ^= num
		}
	}
	return once
}
