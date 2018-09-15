// 695. Max Area of Island

// Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

// Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)
//
// Example 1:
// [[0,0,1,0,0,0,0,1,0,0,0,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,1,1,0,1,0,0,0,0,0,0,0,0],
//  [0,1,0,0,1,1,0,0,1,0,1,0,0],
//  [0,1,0,0,1,1,0,0,1,1,1,0,0],
//  [0,0,0,0,0,0,0,0,0,0,1,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,0,0,0,0,0,0,1,1,0,0,0,0]]
// Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.

// Example 2:
// [[0,0,0,0,0,0,0,0]]
// Given the above grid, return 0.

// Note: The length of each dimension in the given grid does not exceed 50.

package leetcode

func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	var maxArea int
	var dfs func([][]int, [][]bool, int, int) int
	dfs = func(grid [][]int, visited [][]bool, i, j int) int {
		if grid[i][j] == 0 || visited[i][j] {
			return 0
		}

		visited[i][j] = true

		area := 1
		if i > 0 {
			area += dfs(grid, visited, i-1, j)
		}
		if i+1 < len(grid) {
			area += dfs(grid, visited, i+1, j)
		}
		if j > 0 {
			area += dfs(grid, visited, i, j-1)
		}
		if j+1 < len(grid[i]) {
			area += dfs(grid, visited, i, j+1)
		}

		if area > maxArea {
			maxArea = area
		}
		return area
	}

	for i := range grid {
		for j := range grid[i] {
			dfs(grid, visited, i, j)
		}
	}
	return maxArea
}
