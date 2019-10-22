package main

//完全背包,即物品可以重复拿,用dp累加的形式
//dp[j]已经能组合的个数,dp[j-coin]加上当前coin又可以组合成的个数
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}
