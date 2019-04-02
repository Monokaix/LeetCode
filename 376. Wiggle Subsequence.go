package main

//动态规划，但是可以优化空间复杂度，只使用一个变量即可
//一升一降，所以是up时，down+1，down时，up+1
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up := 1
	down := 1
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			up = down + 1
		} else if nums[i] > nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}
