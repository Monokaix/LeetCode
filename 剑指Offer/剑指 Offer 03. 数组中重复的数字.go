package main

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	res := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			res = nums[i]
			break
		} else {
			m[nums[i]] = true
		}
	}

	return res
}
