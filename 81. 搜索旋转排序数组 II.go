package main

func search2(nums []int, target int) bool {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := (begin + end) / 2
		if nums[mid] == target {
			return true
		} else if nums[begin] < nums[mid] { //左侧有序
			if target >= nums[begin] && target < nums[mid] { //在有序区间范围内
				end = mid - 1
			} else {
				begin = mid + 1
			}
		} else {
			//注意此时要注意重复情况,因为有重复的,所以不用担心漏掉某个值
			if nums[begin] == nums[mid] {
				begin++
			} else { //右侧有序
				if target > nums[mid] && target <= nums[end] {
					begin = mid + 1 //在有序区间范围内
				} else {
					end = mid - 1 //继续寻找一个有序的区间
				}
			}

		}
	}
	return false
}
