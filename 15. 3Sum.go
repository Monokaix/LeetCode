package main

import (
	"sort"
	"fmt"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)
	if len(nums) < 3 || nums[0] > 0 {
		return nil
	}
	for k := 0; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		target := 0-nums[k]
		i := k+1
		j := len(nums)-1
		for i < j{
			if nums[i] + nums[j] == target{
				slice := make([]int,3)
				slice[0] = nums[k]
				slice[1] = nums[i]
				slice[2] = nums[j]
				res = append(res, slice)
				for i < j && nums[i] == nums[i+1]{
					i++
				}

				for i < j && nums[j] == nums[j-1]{
					j--
				}
				i++
				j--
			}else if nums[i] + nums[j] < target{
				i++
			}else{
				j--
			}
		}

	}
	return res
}
