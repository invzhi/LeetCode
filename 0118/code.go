// 118. Pascal's Triangle

// Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it.

// Example:
// Input: 5
// Output:
// [
//      [1],
//     [1,1],
//    [1,2,1],
//   [1,3,3,1],
//  [1,4,6,4,1]
// ]

package leetcode

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for r := 0; r < numRows; r++ {
		triangle[r] = make([]int, r+1)
		triangle[r][0] = 1
		for c := 1; c < r; c++ {
			triangle[r][c] = triangle[r-1][c-1] + triangle[r-1][c]
		}
		triangle[r][r] = 1
	}
	return triangle
}
