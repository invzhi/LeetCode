// 231. Power of Two

// Given an integer, write a function to determine if it is a power of two.

// Example 1:
// Input: 1
// Output: true
// Explanation: 20 = 1

// Example 2:
// Input: 16
// Output: true
// Explanation: 24 = 16

// Example 3:
// Input: 218
// Output: false

package leetcode

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
