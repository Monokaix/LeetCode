package main

import "sort"

// 贪心算法，删除最少，即保留最多，先将interval排序，结尾值越小越靠前，
// 这样留给后面可以填充的interval个数就越大
type Interval struct {
	Start int
	End   int
}
type intervalList []Interval

func (p intervalList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p intervalList) Len() int      { return len(p) }
func (p intervalList) Less(i, j int) bool {
	if p[i].End != p[j].End {
		return p[i].End < p[j].End
	}
	return p[i].Start < p[j].Start
}
func eraseOverlapIntervals(intervals []Interval) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	res := 1
	pre := 0
	sort.Sort(intervalList(intervals))
	//从1开始遍历，因为默认1已经加入
	for i := 1; i < n; i++ {
		if intervals[i].Start >= intervals[pre].End {
			res ++
			pre = i
		}
	}
	return len(intervals) - res
}
