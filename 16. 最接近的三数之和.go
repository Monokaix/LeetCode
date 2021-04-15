package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)

	res := 0
	mindiff := math.MaxInt32
	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1
		for left < right {
			cur := nums[i] + nums[left] + nums[right]
			curDiff := abs(cur, target)

			if curDiff == 0 {
				return cur
			}

			if curDiff < mindiff {
				res = cur
				mindiff = curDiff
			}
			if cur < target {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
