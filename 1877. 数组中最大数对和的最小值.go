package main

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i, j := 0, len(nums)-1; i < j; {
		res = max(res, nums[i]+nums[j])
		i++
		j--
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
