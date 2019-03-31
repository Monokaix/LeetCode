package main

import "math"

//动态规划
//第i天的最大收益等于前i-1天的最大收益或者第i天收益减去前i-1天中的最小收益
func maxProfit(prices []int) int {
	min := math.MaxInt32
	max := 0
	for i := 0; i < len(prices); i++ {
		min = Min(min, prices[i])
		max = Max(max, prices[i]-min)
	}
	return max
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
