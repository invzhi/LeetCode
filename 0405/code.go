// 405. Convert a Number to Hexadecimal

// Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.

// Note:
// 1. All letters in hexadecimal (a-f) must be in lowercase.
// 2. The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.
// 3. The given number is guaranteed to fit within the range of a 32-bit signed integer.
// 4. You must not use any method provided by the library which converts/formats the number to hex directly.

// Example 1:
// Input:
// 26
// Output:
// "1a"

// Example 2:
// Input:
// -1
// Output:
// "ffffffff"

package leetcode

func toHex(num int) string {
	var b []byte
	for i := 32; i > 0; i -= 4 {
		hex := num >> uint(i-4) & 0xF
		if b == nil && hex > 0 {
			b = make([]byte, 0, i/4)
		}
		if b != nil {
			b = append(b, toByte(hex))
		}
	}
	if b == nil {
		return "0"
	}
	return string(b)
}

func toByte(num int) byte {
	if num < 10 {
		return '0' + byte(num)
	}
	return 'a' + byte(num-10)
}
