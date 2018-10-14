// 709. To Lower Case

// Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

// Example 1:
// Input: "Hello"
// Output: "hello"

// Example 2:
// Input: "here"
// Output: "here"

// Example 3:
// Input: "LOVELY"
// Output: "lovely"

package leetcode

func toLowerCase(str string) string {
	b := []byte(str)
	const d = 'a' - 'A'
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += d
		}
	}
	return string(b)
}
