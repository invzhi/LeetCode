// 43. Multiply Strings

// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

// Example 1:
// Input: num1 = "2", num2 = "3"
// Output: "6"

// Example 2:
// Input: num1 = "123", num2 = "456"
// Output: "56088"

// Note:
// 1. The length of both num1 and num2 is < 110.
// 2. Both num1 and num2 contain only digits 0-9.
// 3. Both num1 and num2 do not contain any leading zero, except the number 0 itself.
// 4. You must not use any built-in BigInteger library or convert the inputs to integer directly.

package leetcode

import "strings"

func multiply(num1, num2 string) string {
	ints := make([]int, len(num1)+len(num2))
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			ints[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}

	for i := len(ints) - 1; i > 0; i-- {
		ints[i-1] += ints[i] / 10
		ints[i] %= 10
	}
	return intsToString(ints)
}

func intsToString(ints []int) string {
	var i int
	for i < len(ints)-1 && ints[i] == 0 {
		i++
	}
	ints = ints[i:]

	var b strings.Builder
	b.Grow(len(ints))
	for _, i := range ints {
		b.WriteByte('0' + byte(i))
	}
	return b.String()
}
