package main

import "math"

//思路：dp 注意题目要求连续！！！
//因为要求是连续，所以用dp[i]表示以第i个数字结尾所能表示的最大连续和，最后再比较数组得到最大值
//当dp[i-1]<0时，则dp[i]即为nums[i]，否则dp[i]即为dp[i-1]+nums[i]
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = int(math.Max(float64(nums[i]), float64(dp[i-1]+nums[i])))
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
