// 22. Generate Parentheses

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// For example, given n = 3, a solution set is:
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

package leetcode

func generateParenthesis(n int) []string {
	pair := make([]byte, n+n)
	return dfs(nil, pair, n, 0, 0)
}

func dfs(pairs []string, pair []byte, n, left, right int) []string {
	if left == n && right == n {
		return append(pairs, string(pair))
	}

	if left < n {
		pair[left+right] = '('
		pairs = dfs(pairs, pair, n, left+1, right)
	}
	if right < left {
		pair[left+right] = ')'
		pairs = dfs(pairs, pair, n, left, right+1)
	}
	return pairs
}
