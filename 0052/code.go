// 52. N-Queens II

// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
// Given an integer n, return the number of distinct solutions to the n-queens puzzle.

// Example:
// Input: 4
// Output: 2
// Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
// [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],
//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]

package leetcode

func totalNQueens(n int) int {
	return dfs(nil, n)
}

func dfs(xs []int, n int) int {
	qy := len(xs)
	if qy == n {
		return 1
	}

	var cnt int
qx:
	for qx := 0; qx < n; qx++ {
		for y, x := range xs {
			if x == qx || qy-y == x-qx || qy-y == qx-x {
				continue qx
			}
		}
		cnt += dfs(append(xs, qx), n)
	}
	return cnt
}
