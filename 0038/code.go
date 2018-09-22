// 38. Count and Say

// The count-and-say sequence is the sequence of integers with the first five terms as following:
// 1. 1
// 2. 11
// 3. 21
// 4. 1211
// 5. 111221

// 1 is read off as "one 1" or 11.
// 11 is read off as "two 1s" or 21.
// 21 is read off as "one 2, then one 1" or 1211.

// Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.
// Note: Each term of the sequence of integers will be represented as a string.

// Example 1:
// Input: 1
// Output: "1"

// Example 2:
// Input: 4
// Output: "1211"

package leetcode

func countAndSay(n int) string {
	b := []byte{'1'}
	for n > 1 {
		b = say(b)
		n--
	}
	return string(b)
}

func say(b []byte) []byte {
	var buf []byte
	for i, j := 0, 1; i < len(b); i = j {
		for j < len(b) && b[i] == b[j] {
			j++
		}
		buf = append(buf, byte(j-i+'0'), b[i])
	}
	return buf
}
