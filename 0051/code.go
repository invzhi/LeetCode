// 51. N-Queens

// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle.
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

// Example:
// Input: 4
// Output: [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],
//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.

package leetcode

func solveNQueens(n int) [][]string {
	return dfs(nil, nil, n)
}

func dfs(solutions [][]string, xs []int, n int) [][]string {
	qy := len(xs)
	if qy == n {
		return append(solutions, solution(xs, n))
	}

qx:
	for qx := 0; qx < n; qx++ {
		for y, x := range xs {
			if x == qx || qy-y == x-qx || qy-y == qx-x {
				continue qx
			}
		}
		solutions = dfs(solutions, append(xs, qx), n)
	}
	return solutions
}

func solution(xs []int, n int) []string {
	b := make([]byte, n)
	for i := range b {
		b[i] = '.'
	}

	s := make([]string, n)
	for y, x := range xs {
		b[x] = 'Q'
		s[y] = string(b)
		b[x] = '.'
	}
	return s
}
