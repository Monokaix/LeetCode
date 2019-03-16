package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	cur := []int{}
	combination4(k, n, cur, &res, 1)
	return res
}

func combination4(k int, n int, cur []int, res *[][]int, begin int) {
	if len(cur) == k {
		if n == 0 {
			temp := make([]int, len(cur))
			copy(temp, cur)
			*res = append(*res, temp)
			return
		} else if n < 0 {
			return
		}

	}
	for i := begin; i <= 9; i++ {
		cur = append(cur, i)
		combination4(k, n-i, cur, res, i+1) //i+1，因为不可以包括自身，对比problem39
		cur = cur[:len(cur)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
