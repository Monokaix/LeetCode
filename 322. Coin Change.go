package main

import (
	"math"
	"fmt"
)

//动态规划，类似于背包，一个物品考虑放或者不放，可以重复放
func coinChange(coins []int, amount int) int {
	n := len(coins)
	if n == 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1) //dp[i]表示硬币能组成的和为i的最小个数
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		for j := 0; j < n; j++ {
			if i-coins[j] >= 0 { //因为数组不是有序的，所以如果一起判断if，下面的coins[j]就无法继续执行了，所以有的解法先排序了
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coins[j]]+1)))
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	fmt.Println(dp)
	return dp[amount]
}
//dp[i]状态变化
//coins=[1, 2, 5]  amount=11
//0 1 1 2 2 1 2 2 3 3 2 3 amount
func main(){
	coinChange([]int{1,2,5},11)
}