package main

func sortArray(nums []int) []int {
	quickSort2(0, len(nums)-1, nums)
	return nums
}

func partition4(i, j int, nums []int) int {
	key := nums[i]
	left := i
	for i < j {
		for i < j && nums[j] > key {
			j--
		}
		// 注意先后顺序不能颠倒
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

func quickSort2(i, j int, nums []int) {
	if i >= j {
		return
	}
	index := partition4(i, j, nums)
	quickSort2(i, index-1, nums)
	quickSort2(index+1, j, nums)
}
