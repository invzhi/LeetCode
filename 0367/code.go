// 367. Valid Perfect Square

// Given a positive integer num, write a function which returns True if num is a perfect square else False.
// Note: Do not use any built-in library function such as sqrt.

// Example 1:
// Input: 16
// Output: true

// Example 2:
// Input: 14
// Output: false

package leetcode

func isPerfectSquare(num int) bool {
	lo, hi := 0, num
	for lo <= hi {
		mid := lo + (hi-lo)/2

		m2 := mid * mid
		switch {
		case m2 < num:
			lo = mid + 1
		case m2 > num:
			hi = mid - 1
		default:
			return true
		}
	}
	return false
}
