package main

import "sort"
//贪心算法，先排序，再分发
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i := len(g)-1
	j := len(s)-1
	res := 0
	for i >=0 && j >=0{
		if s[j] >=g[i]{
			res++
			i--
			j--
		}else{
			i--
		}
	}
	return res
}