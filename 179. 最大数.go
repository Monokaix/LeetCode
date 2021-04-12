package main

import (
	"sort"
	"strconv"
)

// 对于两个数，a、b，如果组合ab大于组合ba，则把a放前面，因此
// 将能组合出最大的数字放前面即可，也就是根据这个规则进行排序
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		basex, basey := 10, 10
		for basex <= nums[j] {
			basex *= 10
		}
		for basey <= nums[i] {
			basey *= 10
		}
		return nums[i]*basex+nums[j] > nums[j]*basey+nums[i]
	})

	n := len(nums)
	if n == 0 || nums[0] == 0 {
		return "0"
	}

	res := ""
	for i := 0; i < n; i++ {
		res += strconv.Itoa(nums[i])
	}

	return res
}
