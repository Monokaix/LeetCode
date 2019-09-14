package main

//思路:计算每个位置可以得到的雨水数,对与索引位置i,其左边0-i-1位置的最大值为maxLeft,右边最大值为maxRight
//则i位置能接的雨水数为两者之间的最小值,最后累加即可
//方法一:用两个数组存左边的最大值和右边最大值,需要两个数组存放,额外空间O(2n)
//方法二:用两个变量存左边和右边最大值,滑动窗口
func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}
	value := 0
	leftMax := height[0]
	rightMax := height[n-1]
	l := 1     //左边指针从第一个位置开始
	r := n - 1 //右边指针从倒数第二个位置开始
	for l <= r {
		if leftMax <= rightMax { //结算左边
			cur := leftMax - height[l]
			if cur > 0 {
				value += cur
			}
			if height[l] > leftMax { //更新左边最大值
				leftMax = height[l]
			}
			l++
		} else { //结算右边
			cur := rightMax - height[r]
			if cur > 0 {
				value += cur
			}
			if height[r] > rightMax { //更新右边最大值
				rightMax = height[r]
			}
			r--
		}
	}
	return value
}
