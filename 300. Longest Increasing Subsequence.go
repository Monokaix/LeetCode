package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 //初始每个元素最长子序列都为1
	}
	//dp[i]记录以第i个位置结束的最长子序列长度
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i], 1+dp[j])
			}

		}
	}
	res := 1
	for i := 0; i < n; i++ {
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
