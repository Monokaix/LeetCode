package main

//贪心策略：既然要跳的次数最少，则最开始一直不跳，直到没办法不得不跳时再跳，此时开始跳，cnt++
//此时跳跃到最远的地方，因此要用一个pre记录最大的跳跃位置
//当跳完之后，则更新当前可以跳跃到的最远位置
func jump(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	current := nums[0] // 当前可以到达的位置
	pre := nums[0]     // 遍历过程中，可以到达的最大位置
	cnt := 1
	for i := 1; i < n; i++ {
		if i > current { //需要跳跃，因为贪心到达最后
			cnt++
			current = pre // 同时更新当前可以到达的最远位置
		}
		if pre < nums[i]+i {
			pre = nums[i] + i
		}
	}
	return cnt
}
