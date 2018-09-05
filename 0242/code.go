// 242. Valid Anagram

// Given two strings s and t , write a function to determine if t is an anagram of s.

// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:
// Input: s = "rat", t = "car"
// Output: false

// Note:
// You may assume the string contains only lowercase alphabets.

// Follow up:
// What if the inputs contain unicode characters? How would you adapt your solution to such case?

package leetcode

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, r := range s {
		m[r]++
	}

	for _, r := range t {
		if m[r] == 0 {
			return false
		}
		m[r]--
	}

	return true
}
