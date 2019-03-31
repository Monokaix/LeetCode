package main

import "math"

//两次dp
//一次考虑从0-n-2，因为到n-1的话就成环了首尾相连也是相邻不满足题意
//一次考虑从1到n-1
func rob3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		if nums[0] < nums[1] {
			return nums[1]
		}
		return nums[0]
	}
	dp1 := make([]int, n-1)
	dp2 := make([]int, n-1)
	dp1[0] = nums[0]
	dp1[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	dp2[0] = nums[1] //因为从下标1开始
	dp2[1] = int(math.Max(float64(nums[1]), float64(nums[2])))
	for i := 2; i < n-1; i++ {
		dp1[i] = int(math.Max(float64(dp1[i-1]), float64(nums[i]+dp1[i-2])))
		dp2[i] = int(math.Max(float64(dp2[i-1]), float64(nums[i+1]+dp2[i-2])))
	}
	return int(math.Max(float64(dp1[n-2]), float64(dp2[n-2])))
}
