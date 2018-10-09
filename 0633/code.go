// 633. Sum of Square Numbers

// Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a*a + b*b = c.

// Example 1:
// Input: 5
// Output: True
// Explanation: 1 * 1 + 2 * 2 = 5

// Example 2:
// Input: 3
// Output: False

package leetcode

func judgeSquareSum(c int) bool {
	for i := 2; i*i <= c; i++ {
		if c%i == 0 {
			var cnt int
			for c%i == 0 {
				c /= i
				cnt++
			}

			if cnt&1 == 1 && i%4 == 3 {
				return false
			}
		}
	}
	return c%4 != 3
}
