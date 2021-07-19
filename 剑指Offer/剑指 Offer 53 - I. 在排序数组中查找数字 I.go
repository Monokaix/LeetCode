package main

import (
	"sort"
)

func search(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	}) - sort.Search(len(nums), func(i int) bool {
		return nums[i] > target-1
	})
}
