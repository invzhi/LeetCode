// 201. Bitwise AND of Numbers Range

// Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.

// Example 1:
// Input: [5,7]
// Output: 4

// Example 2:
// Input: [0,1]
// Output: 0

package leetcode

func rangeBitwiseAnd(m, n int) int {
	diff := m ^ n
	for diff&(diff-1) != 0 {
		diff &= diff - 1
	}
	if diff != 0 {
		diff |= diff - 1
	}
	return m & n & ^diff
}
