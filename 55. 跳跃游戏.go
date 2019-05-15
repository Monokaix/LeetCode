package main

// 贪心：
// 1、先将数组分别加上下标得到每个位置可以跳跃到的位置
// 2、要跳到的位置应尽可能远，因此要在满足约束的条件下，用max_dist记录最大可以跳跃的距离
// 3、约束条件为索引i<=max_dist，因为这样才能保证连续跳跃，即从位置0到当前位置，也即之前的元素可以跳跃到当前的位置
func canJump(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = i + nums[i]
	}
	max_dist := nums[0]
	for i := 0; i < n; i++ {
		if i > max_dist{
			return false
		}
		if max_dist < nums[i]{
			max_dist = nums[i]
			if max_dist >= n-1 { //可以跳跃到最后一个元素，因为是一直保证连续跳跃，所以能跳到max_dist的
				return true      //位置时，也一定可以跳到max_dist之前的所有位置
			}
		}
	}
	return true // 只有一个元素0的情况
}
