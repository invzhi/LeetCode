// 524. Longest Word in Dictionary through Deleting

// Given a string and a string dictionary, find the longest string in the dictionary that can be formed by deleting some characters of the given string. If there are more than one possible results, return the longest word with the smallest lexicographical order. If there is no possible result, return the empty string.

// Example 1:
// Input:
// s = "abpcplea", d = ["ale","apple","monkey","plea"]
// Output:
// "apple"

// Example 2:
// Input:
// s = "abpcplea", d = ["a","b","c"]
// Output:
// "a"

// Note:
// 1. All the strings in the input will only contain lower-case letters.
// 2. The size of the dictionary won't exceed 1,000.
// 3. The length of all the strings in the input won't exceed 1,000.

package leetcode

func findLongestWord(s string, d []string) string {
	var lw string
	for _, w := range d {
		if len(w) < len(lw) || len(w) == len(lw) && w >= lw {
			continue
		}
		var i int
		for j := 0; i < len(w) && j < len(s); j++ {
			if s[j] == w[i] {
				i++
			}
		}
		if i == len(w) {
			lw = w
		}
	}
	return lw
}
