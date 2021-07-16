package main

func minArray(numbers []int) int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		mid := i + (j-i)/2
		if numbers[mid] > numbers[j] { // 说明最小值在mid右侧
			i = mid + 1
		} else if numbers[mid] < numbers[j] { // 说明最小值在mid左侧，且有可能就是mid，所以应该是high=mid
			j = mid
		} else {
			j--
		}
	}

	return numbers[i]
}
