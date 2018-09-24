// 67. Add Binary

// Given two binary strings, return their sum (also a binary string).
// The input strings are both non-empty and contains only characters 1 or 0.

// Example 1:
// Input: a = "11", b = "1"
// Output: "100"

// Example 2:
// Input: a = "1010", b = "1011"
// Output: "10101"

package leetcode

func addBinary(a, b string) string {
	var sum []byte

	var prev byte
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var x, y byte
		if i >= 0 {
			x = a[i] - '0'
		}
		if j >= 0 {
			y = b[j] - '0'
		}

		v := x + y + prev
		prev = v & 2 >> 1
		sum = append(sum, '0'+(v&1))
	}
	if prev == 1 {
		sum = append(sum, '1')
	}

	for i, j := 0, len(sum)-1; i < j; i, j = i+1, j-1 {
		sum[i], sum[j] = sum[j], sum[i]
	}
	return string(sum)
}
