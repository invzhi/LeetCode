// 3. Longest Substring Without Repeating Characters

// Given a string, find the length of the longest substring without repeating characters.

// Example 1:
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

package leetcode

func lengthOfLongestSubstring(s string) int {
	var length int
	index := make(map[byte]int)
	for i, j := 0, 0; j < len(s); j++ {
		if index[s[j]] > i {
			i = index[s[j]]
		}
		if j-i+1 > length {
			length = j - i + 1
		}
		index[s[j]] = j + 1
	}
	return length
}
