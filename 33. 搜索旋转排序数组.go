package main

func search(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := (begin + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[end] { //右侧有序
			if target > nums[mid] && target <= nums[end] {
				begin = mid + 1 //在有序区间范围内
			} else {
				end = mid - 1 //继续寻找一个有序的区间
			}
		} else { //左侧有序
			if target >= nums[begin] && target < nums[mid] { //在有序区间范围内
				end = mid - 1
			} else {
				begin = mid + 1
			}
		}
	}
	return -1
}
