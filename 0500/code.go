// 500. Keyboard Row

// Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.

// Example:
// Input: ["Hello", "Alaska", "Dad", "Peace"]
// Output: ["Alaska", "Dad"]

// Note:
// 1. You may use one character in the keyboard more than once.
// 2. You may assume the input string will only contain letters of alphabet.

package leetcode

var rows [26]int

func init() {
	letters := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for r, s := range letters {
		for i := 0; i < len(s); i++ {
			rows[s[i]-'a'] = r
		}
	}
}

func findWords(words []string) []string {
	var i int
	for _, word := range words {
		if isOneRow(word) {
			words[i] = word
			i++
		}
	}
	return words[:i]
}

func isOneRow(word string) bool {
	for i := 1; i < len(word); i++ {
		if row(word[0]) != row(word[i]) {
			return false
		}
	}
	return true

}

func row(b byte) int {
	if b < 'a' {
		b += 'a' - 'A'
	}
	return rows[b-'a']
}
