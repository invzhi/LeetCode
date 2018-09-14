package leetcode

func floodFill(image [][]int, sr, sc int, newColor int) [][]int {
	if oldColor := image[sr][sc]; oldColor != newColor {
		dfs(image, sr, sc, oldColor, newColor)
	}
	return image
}

func dfs(image [][]int, sr, sc int, oldColor, newColor int) {
	if image[sr][sc] != oldColor {
		return
	}

	image[sr][sc] = newColor

	if sr > 0 {
		dfs(image, sr-1, sc, oldColor, newColor)
	}
	if sr+1 < len(image) {
		dfs(image, sr+1, sc, oldColor, newColor)
	}
	if sc > 0 {
		dfs(image, sr, sc-1, oldColor, newColor)
	}
	if sc+1 < len(image[sr]) {
		dfs(image, sr, sc+1, oldColor, newColor)
	}
}
