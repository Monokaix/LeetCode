package main

import "math"

func integerBreak(n int) int {
	m := make([]int, n+1)
	return breakInteger(n, m)
}

//方法一：自顶向下，记忆化搜索，递归
func breakInteger(n int, m []int) int {

	if n == 1 {
		return 1
	}
	if m[n] != 0 {
		return m[n]
	}
	res := -1

	for i := 1; i <= n-1; i++ { //树的分支，每种情况都枚举一遍
		res = int(math.Max(math.Max(float64(res), float64(i*(n-i))), float64(i*breakInteger(n-i, m))))
	}
	m[n] = res
	return res
}

//方法二：自低向上，动态规划
func integerBreak2(n int) int {
	dp := make([]int, n+1) // dp[n]为数字n分割求得的最大值
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			//dp[i],i直接和(i-j)相乘,以及i乘以已经计算得到的dp[i-j]比较取最大值
			dp[i] = int(math.Max(math.Max(float64(dp[i]), float64(j*(i-j))), float64(j*dp[i-j])))
		}
	}
	return dp[n]
}
