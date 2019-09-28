package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	res := 0
	//dp[i]记录以第i个位置结束的最长子序列长度
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i], 1+dp[j])
			}
		}
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
