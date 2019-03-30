package main

import "math"

//动态规划，dp[i][j]记录从原点[0][0]到[i][j]的最短距离。
//其中第一行和第一列需要单独处理，因为从这里出发最短距离是固定的，直接相加即可
//剩余的部分继续采用dp，因为只能向右或者向下移动，所以从是dp[i-1][j]或者dp[i][j-1]开始累加
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//第一行
			if i == 0 && j > 0 { //在原来矩阵里修改
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 && i > 0 { //第一列
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if i > 0 && j > 0 {
				grid[i][j] = int(math.Min(float64(grid[i-1][j]), float64(grid[i][j-1]))) + grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]
}
