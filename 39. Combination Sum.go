package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	cur := []int{}
	combination2(candidates, target, cur, &res, 0)
	return res
}

func combination2(candidates []int, target int, cur []int, res *[][]int, begin int) {
	if target == 0 {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	} else if target < 0 {
		return
	}
	for i := begin; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		combination2(candidates, target-candidates[i], cur, res, i) //i，因为可以包括自身，对比problem77
		cur = cur[:len(cur)-1]
	}
}

func main() {
	num := []int{2, 3, 5}
	fmt.Println(combinationSum(num, 8))
}
