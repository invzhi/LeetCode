// 345. Reverse Vowels of a String

// Write a function that takes a string as input and reverse only the vowels of a string.

// Example 1:
// Input: "hello"
// Output: "holle"

// Example 2:
// Input: "leetcode"
// Output: "leotcede"

// Note: The vowels does not include the letter "y".

package leetcode

func reverseVowels(s string) string {
	b := []byte(s)
	i, j := -1, len(b)
	for {
		for i+1 < len(b) && !isVowel(b[i+1]) {
			i++
		}
		i++
		for j-1 > i && !isVowel(b[j-1]) {
			j--
		}
		j--

		if i >= j {
			break
		}
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
