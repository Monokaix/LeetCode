package main

import (
	"sort"
)

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
	ordered := append(sort.IntSlice{}, nums1...)
	sort.Ints(ordered)
	res := 0
	maxn := 0
	for i := 0; i < n; i++ {
		diff := abs(nums1[i], nums2[i])
		j := sort.SearchInts(ordered, nums2[i])
		if j < n {
			maxn = max(maxn, diff-(ordered[j]-nums2[i]))
		}
		if j > 0 {
			maxn = max(maxn, diff-(nums2[i]-ordered[j-1]))
		}
		res += diff
	}

	return (res - maxn) % (1e9 + 7)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
