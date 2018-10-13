// 541. Reverse String II

// Given a string and an integer k, you need to reverse the first k characters for every 2k characters counting from the start of the string. If there are less than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and left the other as original.

// Example:
// Input: s = "abcdefg", k = 2
// Output: "bacdfeg"

// Restrictions:
// 1. The string consists of lower English letters only.
// 2. Length of the given string and k will in the range [1, 10000]

package leetcode

func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(b); i += k + k {
		r := i + k - 1
		if r >= len(b) {
			r = len(b) - 1
		}

		for l := i; l < r; l, r = l+1, r-1 {
			b[l], b[r] = b[r], b[l]
		}
	}
	return string(b)
}
