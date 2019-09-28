package main

import (
	"sort"
)

//若w和h都不相等,则直接先按w从小到大排序,然后在排完序的h上找最长上升子序列
//但w和h可能相等,而规定w相等时装不进去,因此按w从小到大排序,w相等时h从大到小排序,这样
//当w相等时不用考虑后面的h比当前h大的情况,因为h是从大到小排序,后面的h不会对当前w相等时的h造成干扰
type Pair2 struct {
	w int
	h int
}
type PairList []Pair2

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int      { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].w == p[j].w {
		return p[i].h > p[j].h
	}
	return p[i].w < p[j].w
}
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}
	p := make(PairList, n)
	for i := 0; i < n; i++ {
		pair := Pair2{
			w: envelopes[i][0],
			h: envelopes[i][1],
		}
		p[i] = pair
	}
	sort.Sort(p) //实现对p的排序
	res := 0
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if p[j].h < p[i].h {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res + 1
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
