// 326. Power of Three

// Given an integer, write a function to determine if it is a power of three.

// Example 1:
// Input: 27
// Output: true

// Example 2:
// Input: 0
// Output: false

// Example 3:
// Input: 9
// Output: true

// Example 4:
// Input: 45
// Output: false

// Follow up:
// Could you do it without using any loop / recursion?

package leetcode

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
