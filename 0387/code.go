// 387. First Unique Character in a String

// Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

// Examples:
// s = "leetcode"
// return 0.
// s = "loveleetcode",
// return 2.

// Note: You may assume the string contain only lowercase letters.

package leetcode

func firstUniqChar(s string) int {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if cnt[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
