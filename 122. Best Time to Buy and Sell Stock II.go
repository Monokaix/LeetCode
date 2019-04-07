package main

func maxProfit3(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			res += tmp
		}
	}
	return res
}
