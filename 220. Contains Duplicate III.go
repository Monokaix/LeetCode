package main

import "math"

//1.是一个滑动窗口＋查查找的问题
//2.用一个map记录大小为k+1大小的滑动窗口
//3.遍历数组如果在map中存在一个范围在(nums[i]-t,nums[i]+t)的数，返回true
//4.若滑动窗口大小超过k+1，删除窗口第一个元素，继续滑动
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := make(map[int]bool)
	for i := 0;i<len(nums);i++{
		//if _,ok := m[nums[i]];ok{
		//	return true
		//}
		//取出map中的key进行比较
		for key := range m{
			if math.Abs(float64(key-nums[i])) <= float64(t){
				return true
			}
		}
		m[nums[i]] = true

		//保证窗口内只有k+1个元素
		if len(m) == k+1{
			delete(m, nums[i-k])
		}
	}
	return false
}

