package main

import "fmt"

var cnt int
var cnt2 int

func lexicalOrder(n int) []int {
	var res []int
	for i := 1; i < 10; i++ {
		dfs6(i, n, &res)
	}
	return res
}

func dfs6(i, n int, res *[]int) {
	cnt2++
	if i > n {
		return
	}
	cnt++
	*res = append(*res, i)
	for j := 0; j < 10; j++ {
		cur := i*10 + j
		dfs6(cur, n, res)
	}
}

func main() {
	res := lexicalOrder(20)
	fmt.Println(res)
	fmt.Println(cnt)
	fmt.Println(cnt2)
}
