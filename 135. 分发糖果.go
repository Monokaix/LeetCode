package main

import "math"

func candy(ratings []int) int {
	n := len(ratings)
	res := make([]int, n)
	//先从左往右遍历使得当前i的数量都比左边多
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}
	//再从右往左遍历使得满足比右边多
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			res[i] = int(math.Max(float64(res[i+1]+1), float64(res[i])))
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += res[i]
	}
	return sum + n
}
