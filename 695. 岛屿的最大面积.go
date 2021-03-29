package main

func main() {

}

var direction2 = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

// 与problem200思路一样，dfs
func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				curMax := spread2(grid, m, n, i, j)
				if curMax > maxArea {
					maxArea = curMax
				}
			}
		}
	}

	return maxArea
}

func spread2(grid [][]int, m, n, row, col int) int {
	if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] != 1 {
		return 0
	}

	size := 1
	grid[row][col] = -1
	for i := 0; i < len(direction2); i++ {
		size += spread2(grid, m, n, row+direction2[i][0], col+direction2[i][1])
	}

	return size
}
