package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	cur := []int{}

	combination3(candidates, target, cur, &res, 0)
	return res
}

func combination3(candidates []int, target int, cur []int, res *[][]int, begin int) {
	if target == 0 {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	} else if target < 0 {
		return
	}
	for i := begin; i < len(candidates); i++ {
		if i > begin && candidates[i] == candidates[i-1] {
			continue
		}
		cur = append(cur, candidates[i])
		combination3(candidates, target-candidates[i], cur, res, i+1) //i+1，因为不可以包括自身，对比problem39
		cur = cur[:len(cur)-1]
	}
}

func main() {
	num := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(num, 8))
}
