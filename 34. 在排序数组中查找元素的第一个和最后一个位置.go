package main

//思路：分别二分查找最左边端点和最右边端点

func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	res[0] = findLeft2(nums, target)
	res[1] = findRight(nums, target)
	return res
}

func findLeft2(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := (begin + end) / 2
		if target == nums[mid] { //找到一个元素
			if mid == 0 || target > nums[mid-1] {
				return mid //找到最左边的索引
			}
			end = mid - 1
		} else if target > nums[mid] {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
func findRight(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := (begin + end) / 2
		if target == nums[mid] { //找到一个元素
			if mid == len(nums)-1 || target < nums[mid+1] {
				return mid //找到最右边的索引
			}
			begin = mid + 1
		} else if target > nums[mid] {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
