package main

import "math"

func maxProduct(nums []int) int {
	max := math.MinInt32
	imax := 1
	imin := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 { //遇到负数就交换,因为这时乘个负数最大值会变最小值
			imax, imin = imin, imax
		}
		imin = int(math.Min(float64(imin*nums[i]), float64(nums[i])))
		imax = int(math.Max(float64(imax*nums[i]), float64(nums[i])))
		if imax > max {
			max = imax
		}
	}
	return max
}
