package main

func combine(n int, k int) [][]int {
	res := [][]int{}
	cur := []int{}
	combination(n, k, 1, cur, &res)
	return res
}

func combination(n int, k int, start int, cur []int, res *[][]int) {
	if k == len(cur) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	for i := start; i <= n-(k-len(cur))+1; i++ { //这里i的取值上限这么定义是因为：每次开始的位置不能重复，即每个数字不能重复使用
		cur = append(cur, i)
		combination(n, k, i+1, cur, res) //i+1,因为不能包括自身，对比problem39
		cur = cur[:len(cur)-1]
	}
	return
}
