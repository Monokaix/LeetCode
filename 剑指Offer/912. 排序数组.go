package main

func sortArray(nums []int) []int {
	quickSort(0, len(nums)-1, nums)
	return nums
}

func partition(i, j int, nums []int) int {
	key := nums[i]
	left := i
	for i < j {
		for i < j && nums[j] > key {
			j--
		}
		for i < j && nums[i] <= key {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func quickSort(i, j int, nums []int) {
	if i >= j {
		return
	}
	index := partition(i, j, nums)
	quickSort(i, index-1, nums)
	quickSort(index+1, j, nums)
}
