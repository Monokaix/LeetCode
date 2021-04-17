package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	r := n - 1

	for r-1 >= 0 && nums[r-1] >= nums[r] {
		r--
	}

	// 不是最后一个排列
	if r != 0 {
		var k = n - 1
		for k > 0 {
			if nums[k] > nums[r-1] {
				break
			}
			k--
		}
		nums[r-1], nums[k] = nums[k], nums[r-1]
	}

	swapBetween(r, n-1, nums)
}

func swapBetween(start, end int, nums []int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	//swapBetween(1, nums)
	fmt.Println(nums)

}
