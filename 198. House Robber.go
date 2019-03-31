package main

import "math"

func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return tryRob(nums, 0, memo)
}

//方法一：记忆化搜索
//memo[i] 记录从(i...n-1)所能偷取的最大价值，也是状态
func tryRob(nums []int, index int, memo []int) int {
	n := len(nums)
	if index >= n {
		return 0
	}
	if memo[index] != -1 {
		return memo[index]
	}
	res := 0
	for i := index; i < n; i++ {
		//状态转移方程
		res = int(math.Max(float64(res), float64(nums[i]+tryRob(nums, i+2, memo))))
	}
	memo[index] = res
	return res
}

//方法二:动态规划
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	memo := make([]int, n)
	memo[n-1] = nums[n-1] //体现动态规划自底向上的思想
	for i := n - 2; i >= 0; i-- {
		//记录memo[i]
		for j := i; j < n; j++ {
			if j+2 < n {
				memo[i] = int(math.Max(float64(memo[i]), float64(nums[j]+memo[j+2])))
			} else {
				memo[i] = int(math.Max(float64(memo[i]), float64(nums[j])))
			}
		}
	}
	return memo[0]
}

//另外一种动态规划思路，从memo[0]开始
//public int rob(int[] nums) {
//	int n = nums.length;
//	if (n <= 1) return n == 0 ? 0: nums[0];
//	int[] memo = new int[n];
//	memo[0] = nums[0];
//	memo[1] = Math.max(nums[0], nums[1]);
//	for (int i = 2; i < n; i++)
//	memo[i] = Math.max(memo[i - 1], nums[i] + memo[i - 2]); //偷或者不偷
//	return memo[n - 1];
//}
