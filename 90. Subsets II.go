package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	cur := []int{}
	combination6(nums, cur, &res, 0)
	return res
}
func combination6(nums []int, cur []int, res *[][]int, begin int) {
	temp := make([]int, len(cur))
	copy(temp, cur)
	*res = append(*res, temp)
	for i := begin; i < len(nums); i++ {
		if i > begin && nums[i] == nums[i-1]{
			//fmt.Println("i",i)
			//fmt.Println("begin",begin)
			//fmt.Println("nums[i]",nums[i])
			//fmt.Println("nums[i-1]",nums[i-1])
			//fmt.Println("cur",cur)
			//fmt.Println("%%%%%%%%%%%")
			continue
		}
		cur = append(cur, nums[i])
		combination6(nums, cur, res, i+1) //i+1，因为不可以包括自身，对比problem39
		cur = cur[:len(cur)-1]
	}
	return
}
func main() {
	nums := []int{1, 1}
	fmt.Println(subsetsWithDup(nums))
}