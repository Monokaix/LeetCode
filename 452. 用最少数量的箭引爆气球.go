package main

import "sort"

//贪心：每次尽可能多的有区间重叠，当下一个区间和当前保存的区间无重叠时，需要增加弓箭手
//每次遍历各个区间时判断要不要更新区间的值

type pair struct {
	first  int
	second int
}
type pairlist []pair

func (p pairlist) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p pairlist) Len() int      { return len(p) }
func (p pairlist) Less(i, j int) bool {
	return p[i].first < p[j].first
}
func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	list := make(pairlist, 0)
	for i := 0; i < n; i++ {
		tmp := pair{points[i][0], points[i][1]}
		list = append(list, tmp)
	}
	sort.Sort(list)
	cnt := 1
	begin := list[0].first //初始区间的左边
	end := list[0].second  //初始区间的右边
	for i := 1; i < n; i++ {
		if list[i].first <= end { //与新来的气球有重叠部分，不用加新的气球
			begin = list[i].first //更新begin
			if end > list[i].second { //同时也要考虑尾端
				end = list[i].second
			}
		} else { //没有重叠区间，需要增加弓箭手
			cnt++
			begin = list[i].first //更新新的区间
			end = list[i].second
		}
	}
	return cnt+begin-begin
}
