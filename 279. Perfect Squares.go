package main

import (
	"container/list"
)

//方法一：图的深广度优先遍历
func numSquares(n int) int {
	l := list.New()
	type info struct {
		num  int
		step int
	}
	l.PushBack(info{n, 0})
	visited := make([]bool, n+1)
	visited[n] = true
	for l.Len() > 0 {
		num := l.Front().Value.(info).num
		step := l.Front().Value.(info).step
		l.Remove(l.Front())

		for i := 0; ; i++ {
			a := num - i*i
			if a < 0 {
				break
			}
			if a == 0 {
				return step + 1
			}
			if !visited[a] {
				l.PushBack(info{a, step + 1})
			}
			visited[a] = true
		}
	}
	return 0
}

//动态规划
func numSquares2(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if dp[i] == 0 || dp[i] > dp[i-j*j]+1 {
				dp[i] = dp[i-j*j] + 1
			}

		}
	}
	return dp[n]
}
