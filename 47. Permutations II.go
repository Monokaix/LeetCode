package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	//sort.Ints(nums)
	res := [][]int{}
	temp := []int{}
	used := make([]bool, len(nums))

	generatePermute2(nums, 0, temp, &res, used)
	return res
}

func generatePermute2(nums []int, index int, temp []int, res *[][]int, used []bool) {
	if index == len(nums) {
		tmp := make([]int,len(temp))
		copy(tmp,temp)
		*res = append(*res, tmp)
		return
	}
	taken := make(map[int]bool,len(nums)-index)
	for i := 0; i < len(nums); i++ {
		if !taken[nums[i]] && !used[i] {
			taken[nums[i]] = true
			temp = append(temp, nums[i])
			used[i] = true
			generatePermute2(nums, index+1, temp, res, used)
			temp = temp[:len(temp)-1]
			used[i] = false
		}
	}
}

func main() {
	nums := []int{2,2,1,1}
	fmt.Println(permuteUnique(nums))
}