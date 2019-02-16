// 205. Isomorphic Strings

// Given two strings s and t, determine if they are isomorphic.
// Two strings are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

// Example 1:
// Input: s = "egg", t = "add"
// Output: true

// Example 2:
// Input: s = "foo", t = "bar"
// Output: false

// Example 3:
// Input: s = "paper", t = "title"
// Output: true

// Note: You may assume both s and t have the same length.

package leetcode

func isIsomorphic(s, t string) bool {
	index1 := make([]int, 256)
	index2 := make([]int, 256)
	for i := 0; i < len(s); i++ {
		if index1[s[i]] != index2[t[i]] {
			return false
		}
		index1[s[i]] = i + 1
		index2[t[i]] = i + 1
	}
	return true
}
