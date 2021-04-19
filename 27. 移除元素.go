package main

//func removeElement(nums []int, val int) int {
//	i := 0
//	for j := 0; j < len(nums); j++ {
//		if nums[j] != val {
//			nums[i] = nums[j]
//			i++
//		}
//	}
//
//	return i + 1
//}

// 方法二
func removeElement(nums []int, val int) int {
	i := 0
	n := len(nums)
	j := n - 1
	for i <= j {
		for j >= 0 && nums[j] == val {
			j--
		}
		for i < n && nums[i] != val {
			i++
		}

		if i < n && j >= 0 && i <= j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return i
}
