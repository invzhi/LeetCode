// 14. Longest Common Prefix

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

// Example 1:
// Input: ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Note:
// All given inputs are in lowercase letters a-z.

package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(prefix) {
			prefix = prefix[:len(strs[i])]
		}

		var j int
		for ; j < len(prefix); j++ {
			if strs[i][j] != prefix[j] {
				break
			}
		}

		prefix = prefix[:j]
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}
