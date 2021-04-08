package main

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1

	for low < high {
		mid := (low + high) / 2
		if nums[mid] < nums[high] { // 说明mid右侧递增，则最小值一定在左侧（可能包括mid本身）
			high = mid
		} else {
			low = mid + 1
		}
	}

	// low和high相遇
	return nums[low]
}
