// 557. Reverse Words in a String III

// Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

// Example 1:
// Input: "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

// Note: In the string, each word is separated by single space and there will not be any extra space in the string.

package leetcode

func reverseWords(s string) string {
	b := []byte(s)
	for i, j := 0, 0; i < len(b); i = j {
		for i < len(b) && b[i] == ' ' {
			i++
		}
		j = i + 1
		for j < len(b) && b[j] != ' ' {
			j++
		}

		for k := j - 1; i < k; i, k = i+1, k-1 {
			b[i], b[k] = b[k], b[i]
		}
	}
	return string(b)
}
