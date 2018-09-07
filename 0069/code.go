// 69. Sqrt(x)

// Implement int sqrt(int x).

// Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

// Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

// Example 1:
// Input: 4
// Output: 2

// Example 2:
// Input: 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.

package leetcode

func mySqrt(x int) int {
	lo, hi := 0, x
	for lo < hi {
		mid := hi - (hi-lo)/2

		m2 := mid * mid
		switch {
		case m2 > x:
			hi = mid - 1
		case m2 < x:
			lo = mid
		default:
			return mid
		}
	}
	return lo
}
