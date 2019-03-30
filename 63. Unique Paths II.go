package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	//for i := 0; i <= m; i++ {
	//	dp = append(dp, make([]int, n+1))
	//}
	//dp[0][1] = 1 //这部是为了让dp[1][1]为1
	//for i := 1; i <= m; i++ {
	//	for j := 1; j <= n; j++ {
	//		if obstacleGrid[i-1][j-1] == 1 {
	//			continue
	//		}
	//		dp[i][j] = dp[i-1][j] + dp[i][j-1]
	//	}
	//}
	//return dp[m][n]
	dp := [][]int{}
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j]
			} else if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

