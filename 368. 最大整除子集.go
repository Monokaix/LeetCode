package main

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	maxIndex := 0
	maxNum := 1
	sort.Ints(nums)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if dp[i] > maxNum {
			maxIndex = i
			maxNum = dp[i]
		}
	}

	res := make([]int, 0)
	for i := maxIndex; i >= 0; i-- {
		if nums[maxIndex]%nums[i] == 0 && dp[i] == maxNum {
			res = append(res, nums[i])
			maxNum--
			maxIndex = i
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
