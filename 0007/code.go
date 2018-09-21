// 7. Reverse Integer

// Given a 32-bit signed integer, reverse digits of an integer.

// Example 1:
// Input: 123
// Output: 321

// Example 2:
// Input: -123
// Output: -321

// Example 3:
// Input: 120
// Output: 21

// Note:
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

package leetcode

import "math"

func reverse(x int) int {
	var y int
	for x != 0 {
		pop := x % 10
		if y > math.MaxInt32/10 || (y == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if x < math.MinInt32/10 || (y == math.MinInt32/10 && pop < -8) {
			return 0
		}

		y = y*10 + pop
		x /= 10
	}
	return y
}
