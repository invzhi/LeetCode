// 784. Letter Case Permutation

// Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.  Return a list of all possible strings we could create.

// Examples:
// Input: S = "a1b2"
// Output: ["a1b2", "a1B2", "A1b2", "A1B2"]

// Input: S = "3z4"
// Output: ["3z4", "3Z4"]

// Input: S = "12345"
// Output: ["12345"]

// Note:
// - S will be a string with length between 1 and 12.
// - S will consist only of letters or digits.

package leetcode

func letterCasePermutation(S string) []string {
	s := []string{S}
	for i := 0; i < len(S); i++ {
		if S[i] < '0' || S[i] > '9' {
			for j := len(s) - 1; j >= 0; j-- {
				b := []byte(s[j])
				letterCase(b, i)
				s = append(s, string(b))
			}
		}
	}
	return s
}

func letterCase(b []byte, i int) {
	const d = 'a' - 'A'
	if b[i] >= 'A' && b[i] <= 'Z' {
		b[i] += d
	} else {
		b[i] -= d
	}
}
