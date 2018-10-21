// 168. Excel Sheet Column Title

// Given a positive integer, return its corresponding column title as appear in an Excel sheet.
// For example:
//     1 -> A
//     2 -> B
//     3 -> C
//     ...
//     26 -> Z
//     27 -> AA
//     28 -> AB
//     ...

// Example 1:
// Input: 1
// Output: "A"

// Example 2:
// Input: 28
// Output: "AB"

// Example 3:
// Input: 701
// Output: "ZY"

package leetcode

func convertToTitle(n int) string {
	var b []byte
	for n > 0 {
		n--
		b = append(b, 'A'+byte(n%26))
		n /= 26
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
