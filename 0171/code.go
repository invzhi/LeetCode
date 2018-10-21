// 171. Excel Sheet Column Number

// Given a column title as appear in an Excel sheet, return its corresponding column number.
// For example:
//     A -> 1
//     B -> 2
//     C -> 3
//     ...
//     Z -> 26
//     AA -> 27
//     AB -> 28
//     ...

// Example 1:
// Input: "A"
// Output: 1

// Example 2:
// Input: "AB"
// Output: 28

// Example 3:
// Input: "ZY"
// Output: 701

package leetcode

func titleToNumber(s string) int {
	var num int
	for i := 0; i < len(s); i++ {
		num = num*26 + int(s[i]-'A'+1)
	}
	return num
}
