// 867. Transpose Matrix

// Given a matrix A, return the transpose of A.
// The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

// Example 1:
// Input: [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[1,4,7],[2,5,8],[3,6,9]]

// Example 2:
// Input: [[1,2,3],[4,5,6]]
// Output: [[1,4],[2,5],[3,6]]

// Note:
// 1. 1 <= A.length <= 1000
// 2. 1 <= A[0].length <= 1000

package leetcode

func transpose(A [][]int) [][]int {
	var T [][]int
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if j >= len(T) {
				T = append(T, nil)
			}
			T[j] = append(T[j], A[i][j])
		}
	}
	return T
}
