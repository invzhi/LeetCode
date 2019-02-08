// 17. Letter Combinations of a Phone Number

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// Example:
// Input: "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

// Note: Although the above answer is in lexicographical order, your answer could be in any order you want.

package leetcode

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	buttons := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var combinations []string
	combination := make([]byte, len(digits))

	var f func(int)
	f = func(i int) {
		if i == len(digits) {
			combinations = append(combinations, string(combination))
			return
		}

		letters := buttons[digits[i]-'2']
		for j := 0; j < len(letters); j++ {
			combination[i] = letters[j]
			f(i + 1)
		}
	}
	f(0)

	return combinations
}
