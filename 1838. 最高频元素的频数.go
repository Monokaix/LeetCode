package main

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 1
	l := 0
	sum := 0
	for r := 1; r < n; r++ {
		sum += (nums[r] - nums[r-1]) * (r - l)
		for sum > k {
			sum -= nums[r] - nums[l]
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}
