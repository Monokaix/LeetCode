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
		if i > begin && nums[i] == nums[i-1] { //排完序后，从第二个元素开始，若他与前一个元素相等，则不加入当前元素
			continue                           //i大于begin，如本例中begin=0，i，表示在第一轮递归条件下（因为begin=0）
		}                                      //后续的递归已经有重复元素，则跳过后续轮的递归，也就是说回溯后的第二个1越过了
		cur = append(cur, nums[i])
		combination6(nums, cur, res, i+1) //i+1，因为不可以包括自身，对比problem39
		cur = cur[:len(cur)-1]
	}
}
func main() {
	nums := []int{1, 1}
	fmt.Println(subsetsWithDup(nums))
}
