package main

import "sort"

//首先按左端点进行排序,这样可以合并的区间都在挨着,然后遍历一遍区间进行合并
type interval struct {
	first  int
	second int
}
type intervalist []interval

func (p intervalist) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p intervalist) Len() int      { return len(p) }
func (p intervalist) Less(i, j int) bool {
	return p[i].first < p[j].first
}
func merge2(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return nil
	}
	res := make([][]int, 0)
	list := make(intervalist, 0)
	for i := 0; i < n; i++ {
		list = append(list, interval{intervals[i][0], intervals[i][1]})
	}
	sort.Sort(list)
	for i := 0; i < n; i++ {
		length := len(res)
		//当前区间与最后一个区间无重叠或者当前res为空,直接append
		if length == 0 || list[i].first > res[length-1][1] {
			res = append(res, []int{list[i].first, list[i].second})
		} else { //有重叠,更新区间
			if res[length-1][1] < list[i].second {
				res[length-1][1] = list[i].second
			}
		}
	}
	return res
}
