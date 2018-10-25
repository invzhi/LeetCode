// 20. Valid Parentheses

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.

// Example 1:
// Input: "()"
// Output: true

// Example 2:
// Input: "()[]{}"
// Output: true

// Example 3:
// Input: "(]"
// Output: false

// Example 4:
// Input: "([)]"
// Output: false

// Example 5:
// Input: "{[]}"
// Output: true

package leetcode

func isValid(s string) bool {
	brackets := map[byte]byte{')': '(', '}': '{', ']': '['}

	stack := make([]byte, len(s))
	var i int
	for j := 0; j < len(s); j++ {
		switch s[j] {
		case '(', '{', '[':
			stack[i] = s[j]
			i++
		default:
			if i == 0 || stack[i-1] != brackets[s[j]] {
				return false
			}
			i--
		}
	}
	return i == 0
}
