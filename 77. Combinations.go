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

	for i := start; i <= n-(k-len(cur))+1; i++ {
		cur = append(cur, i)
		combination(n, k, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
	return
}
