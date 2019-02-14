// 74. Search a 2D Matrix

// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
// - Integers in each row are sorted from left to right.
// - The first integer of each row is greater than the last integer of the previous row.

// Example 1:
// Input:
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 3
// Output: true

// Example 2:
// Input:
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 13
// Output: false

package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	lo, hi := 0, len(matrix)-1
	for lo < hi {
		mid := hi - (hi-lo)/2
		switch {
		case matrix[mid][0] < target:
			lo = mid
		case matrix[mid][0] > target:
			hi = mid - 1
		default:
			return true
		}
	}

	i := lo
	lo, hi = 0, len(matrix[i])-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case matrix[i][mid] < target:
			lo = mid + 1
		case matrix[i][mid] > target:
			hi = mid - 1
		default:
			return true
		}
	}
	return false
}
