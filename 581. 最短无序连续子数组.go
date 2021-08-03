package main

import (
	"fmt"
	"math"
)

func findUnsortedSubarray(nums []int) int {
	l, r := -1, -1
	n := len(nums)
	min := math.MaxInt32
	max := math.MinInt32
	for i := 0; i < n; i++ {
		if nums[i] < max {
			r = i
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	for i := n - 1; i >= 0; i-- {
		if nums[i] > min {
			l = i
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	fmt.Println(l, r)

	if l == r {
		return 0
	}
	return r - l + 1
}

func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	findUnsortedSubarray(nums)
}
