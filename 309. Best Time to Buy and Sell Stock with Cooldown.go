package main

func maxProfit2(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	if n == 2 {
		if prices[1]-prices[0] > 0 {
			return prices[1] - prices[0]
		}
		return 0
	}
	dp1 := make([]int, n)  //未持有股票
	dp2 := make([]int, n)  //持有股票
	dp2[0] = 0 - prices[0] //因为第一天是负收益，没有买卖
	for i := 1; i < n; i++ {
		dp1[i] = Max(dp1[i-1], dp2[i-1]+prices[i]) //当天未持有股票：昨天也未持有股票，或者昨天持有，今天卖出
		tmp := 0
		if i >= 2 {
			tmp = dp1[i-2]
		}
		dp2[i] = Max(dp2[i-1], tmp-prices[i])
	}
	return Max(dp1[n-1], dp2[n-1])
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
