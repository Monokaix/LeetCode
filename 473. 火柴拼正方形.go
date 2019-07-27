package main

import "sort"

//思路：回溯法,每次取一个数尝试往四条边里放,只要还可以往某条边放就一直放
//有三点优化
//1、总和除以4有余数,直接返回false
//2、先把数组排序
//3、每条边放置的总长不能超过总和/4
func makesquare(nums []int) bool {
	n := len(nums)
	if n < 4 {
		return false //少于四条边,直接返回false
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	if sum%4 != 0 {
		return false
	}
	sort.Ints(nums)
	target := sum / 4 //每条边的长度
	var edge [4]int   //每条边当前放置策略下的长度
	return tryPlace(n-1, nums, target, edge)

}

func tryPlace(i int, nums []int, target int, edge [4]int) bool {
	if i < 0 { //递归结束,然后判断四条边是否满足条件
		return edge[0] == target && edge[1] == target && edge[2] == target && edge[3] == target
	}
	//往四条边放
	for j := 0; j < 4; j++ {
		if edge[j]+nums[i] > target { //放不下时,尝试放下一条边
			continue
		}
		edge[j] += nums[i] //放在第j个桶中
		if tryPlace(i-1, nums, target, edge) { //当前放置正确,返回true
			return true
		}
		//若放置失败,进行回溯
		edge[j] -= nums[i]
	}
	return false //如四条边都无法放置,说明前面i次放置错误,返回false以便继续进行回溯
}
