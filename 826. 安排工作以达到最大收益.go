package main

import "sort"

type Pair3 struct {
	difficulty int
	profit     int
}

type pairList3 []Pair3

func (p pairList3) Len() int {
	return len(p)
}

func (p pairList3) Less(i, j int) bool {
	if p[i].difficulty == p[j].difficulty {
		return p[i].profit < p[j].profit
	}

	return p[i].difficulty < p[j].difficulty
}

func (p pairList3) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	p := pairList3{}
	n := len(difficulty)
	for i := 0; i < n; i++ {
		p = append(p, Pair3{difficulty[i], profit[i]})
	}
	sort.Sort(p)
	sort.Ints(worker)

	res := 0
	i := 0
	best := 0
	for _, val := range worker {
		for i < n && val >= p[i].difficulty {
			// 这里是按顺序排列的，所以后续工人的最大收益是上一个工人收益best和接下来要比较的收益的最大值
			best = max(best, p[i].profit)
			i++
		}
		res += best
	}

	return res
}
