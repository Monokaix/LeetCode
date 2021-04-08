package main

func findMin2(nums []int) int {
	low := 0
	high := len(nums) - 1

	for low < high {
		mid := (low + high) / 2
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high-- // 相等时无法判断最小值再左侧还是右侧，因为此时mid和high位置想等，因此high--并不会漏掉最小值
		}
	}

	return nums[low]
}
