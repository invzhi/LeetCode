// 37. Sudoku Solver

// Write a program to solve a Sudoku puzzle by filling the empty cells.
// A sudoku solution must satisfy all of the following rules:
// 1. Each of the digits 1-9 must occur exactly once in each row.
// 2. Each of the digits 1-9 must occur exactly once in each column.
// 3. Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
// Empty cells are indicated by the character '.'.

// Note:
// - The given board contain only digits 1-9 and the character '.'.
// - You may assume that the given Sudoku puzzle will have a single unique solution.
// - The given board size is always 9x9.

package leetcode

func solveSudoku(board [][]byte) {
	var row, col, box [9][9]bool
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			cell := board[r][c] - '1'
			row[r][cell] = true
			col[c][cell] = true
			box[r/3*3+c/3][cell] = true
		}
	}
	dfs(board, row[:], col[:], box[:], 0)
}

func dfs(board [][]byte, row, col, box [][9]bool, index int) bool {
	for i := index; i < 81; i++ {
		r, c := i/9, i%9
		if board[r][c] != '.' {
			continue
		}
		for cell := byte(0); cell < 9; cell++ {
			if row[r][cell] || col[c][cell] || box[r/3*3+c/3][cell] {
				continue
			}

			board[r][c] = '1' + cell
			row[r][cell] = true
			col[c][cell] = true
			box[r/3*3+c/3][cell] = true

			if dfs(board, row, col, box, i+1) {
				return true
			}

			board[r][c] = '.'
			row[r][cell] = false
			col[c][cell] = false
			box[r/3*3+c/3][cell] = false
		}
		return false
	}
	return true
}
