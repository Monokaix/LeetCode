package main

//1.是一个滑动窗口＋查查找的问题
//2.用一个map记录大小为k+1大小的滑动窗口
//3.遍历数组如果在map中存在相同元素，返回true
//4.若滑动窗口大小超过k+1，删除窗口第一个元素，继续滑动
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	for i := 0;i<len(nums);i++{
		if _,ok := m[nums[i]];ok{
			return true
		}
		m[nums[i]] = true

		//保证窗口内只有k+1个元素
		if len(m) == k+1{
			delete(m, nums[i-k])
		}
	}
	return false
}
