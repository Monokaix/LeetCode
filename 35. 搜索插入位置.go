package main

func searchInsert(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for {
		mid := (begin + end) / 2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			if mid == len(nums)-1{
				return len(nums)
			}
			if target < nums[mid+1] {
				return mid + 1
			}
			begin = mid + 1
		} else {
			if mid == 0{
				return 0
			}
			if target > nums[mid-1] {
				return mid
			}
			end = mid - 1
		}
	}
}
