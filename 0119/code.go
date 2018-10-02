// 119. Pascal's Triangle II

// Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
// Note that the row index starts from 0.
// In Pascal's triangle, each number is the sum of the two numbers directly above it.

// Example:
// Input: 3
// Output: [1,3,3,1]

// Follow up:
// Could you optimize your algorithm to use only O(k) extra space?

package leetcode

func getRow(rowIndex int) []int {
	var rows [2][]int
	rows[0] = make([]int, rowIndex+1)
	rows[1] = make([]int, rowIndex+1)
	for r := 0; r <= rowIndex; r++ {
		i := r & 1
		rows[i][0] = 1
		for c := 1; c < r; c++ {
			rows[i][c] = rows[i^1][c-1] + rows[i^1][c]
		}
		rows[i][r] = 1
	}
	return rows[rowIndex&1]
}
