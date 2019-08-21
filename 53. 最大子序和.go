package main

//状态转移方程 Sum[i] = max(a[i]+Sum[i-1],a[i]),curSum[i]表示以i结尾的求出的最大值
//表示最大值要么为当前已求得的curSum加上当前元素a[i]后的值，要么是a[i],因为负数加个数永远不会有这个数本身大(从a[i]开始重新累加,curSum清零)
func maxArrSeq(nums []int) int {
	maxSum := nums[0]
	curSum := 0
	for i := 0; i < len(nums); i++ {
		if curSum > 0 {
			curSum += nums[i] //a[i]+Sum[i-1]
		} else {
			curSum = nums[i] //a[i]
		}
		//找出最大值
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
