// 137. Single Number II

// Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.
// Note:
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

// Example 1:
// Input: [2,2,3,2]
// Output: 3

// Example 2:
// Input: [0,1,0,1,0,1,99]
// Output: 99

package leetcode

func singleNumber(nums []int) int {
	var one, two, three int

	for _, num := range nums {
		two |= one & num
		one ^= num

		three = one & two
		one &= ^three
		two &= ^three
	}
	return one
}
