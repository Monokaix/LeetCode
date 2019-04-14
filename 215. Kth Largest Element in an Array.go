package main

func findKthLargest(nums []int, k int) int {
	start := 0
	end := len(nums) - 1
	kth := end - k + 1
	index := partition3(nums, start, end)
	for index != kth {
		if index > kth {
			end = index - 1
			index = partition3(nums, start, end) // 往左边找
		} else
		{
			start = index + 1
			index = partition3(nums, start, end) // 往右边找
		}
	}
	return nums[index]
}
func partition3(nums []int, left int, right int) int {
	for left < right {
		for left < right && nums[left] <= nums[right] {
			right--
		}
		//开始交换
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++ //刚交换过，左边继续往右走一步
		}

		for left < right && nums[left] <= nums[right] {
			left++
		}

		//开始交换
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			right-- //刚交换过，右边继续往左走一步
		}
	}
	return left
}
