package main

import "math"

func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return tryRob(nums, 0, memo)
}

//方法一：记忆化搜索
//memo[i] 记录从(i...n-1)所能偷取的最大价值，也是状态
func tryRob(nums []int, index int, memo []int) int {
	n := len(nums)
	if index >= n {
		return 0
	}
	if memo[index] != -1 {
		return memo[index]
	}
	res := 0
	for i := index; i < n; i++ {
		//状态转移方程
		res = int(math.Max(float64(res), float64(nums[i]+tryRob(nums, i+2, memo))))
	}
	memo[index] = res
	return res
}

//动态规划思路，从dp[0]开始
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n < 2 {
		return nums[0]
	}
	dp := make([]int, n) //dp[i]表示到第i个房子为止可以偷窃的最大值
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < n; i++ {
		//dp[i-1]表示不偷第i个,nums[i]+dp[i-2]表示偷第i个
		dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i]+dp[i-2])))
	}
	return dp[n-1]
}
