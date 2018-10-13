// 344. Reverse String

// Write a function that takes a string as input and returns the string reversed.

// Example 1:
// Input: "hello"
// Output: "olleh"

// Example 2:
// Input: "A man, a plan, a canal: Panama"
// Output: "amanaP :lanac a ,nalp a ,nam A"

package leetcode

func reverseString(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
