// 415. Add Strings

// Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.
// Note:
// 1. The length of both num1 and num2 is < 5100.
// 2. Both num1 and num2 contains only digits 0-9.
// 3. Both num1 and num2 does not contain any leading zero.
// 4. You must not use any built-in BigInteger library or convert the inputs to integer directly.

package leetcode

func addStrings(num1, num2 string) string {
	var b []byte

	var k byte
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := k
		if i >= 0 {
			sum += num1[i] - '0'
		}
		if j >= 0 {
			sum += num2[j] - '0'
		}
		k = sum / 10
		b = append(b, '0'+sum%10)
	}
	if k > 0 {
		b = append(b, '1')
	}

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
