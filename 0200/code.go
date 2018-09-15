// 200. Number of Islands

// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

// Example 1:
// Input:
// 11110
// 11010
// 11000
// 00000
// Output: 1
//
// Example 2:
// Input:
// 11000
// 11000
// 00100
// 00011
// Output: 3

package leetcode

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	var num int
	for i, r := range grid {
		for j, c := range r {
			if c == '0' || visited[i][j] {
				continue
			}

			num++
			visit(grid, visited, i, j)
		}
	}
	return num
}

func visit(grid [][]byte, visited [][]bool, i, j int) {
	if grid[i][j] == '0' || visited[i][j] {
		return
	}

	visited[i][j] = true

	if i > 0 {
		visit(grid, visited, i-1, j)
	}
	if i+1 < len(grid) {
		visit(grid, visited, i+1, j)
	}
	if j > 0 {
		visit(grid, visited, i, j-1)
	}
	if j+1 < len(grid[i]) {
		visit(grid, visited, i, j+1)
	}
}
