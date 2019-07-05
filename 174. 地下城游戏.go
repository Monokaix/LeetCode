package main

import "math"

//dp 从右下角往上推，注意处理最后一行和最后一列，类似于64号问题
//dp[i][j]表示在i j位置需要的最小血量
func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 { //终点
				dp[i][j] = int(math.Max(float64(1), float64(1-dungeon[i][j])))
			} else if i == m-1 { //最后一行
				dp[i][j] = int(math.Max(float64(1), float64(dp[i][j+1]-dungeon[i][j])))
			} else if j == n-1 { //最后一列
				dp[i][j] = int(math.Max(float64(1), float64(dp[i+1][j]-dungeon[i][j])))
			} else {
				dp[i][j] = int(math.Max(float64(1), math.Min(float64(dp[i+1][j]), float64(dp[i][j+1]))-float64(dungeon[i][j])))
			}
		}
	}
	return dp[0][0]
}
