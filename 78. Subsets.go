package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	combination5(nums, cur, &res, 0)
	return res
}

func combination5(nums []int, cur []int, res *[][]int, begin int) {
	temp := make([]int, len(cur))
	copy(temp, cur)
	*res = append(*res, temp) //这里第一次加入的是一个空集
	for i := begin; i < len(nums); i++ {
		cur = append(cur, nums[i])
		combination5(nums, cur, res, i+1) //i+1，因为不可以包括自身，对比problem39
		cur = cur[:len(cur)-1]
	}
	return
}
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
