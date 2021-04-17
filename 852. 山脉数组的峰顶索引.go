package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	i := 0
	n := len(arr)
	j := n - 1

	for i < j {
		mid := i + (j-i)/2

		if arr[mid] < arr[mid+1] {
			i = mid + 1
		} else {
			j = mid
		}
	}

	return i
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(peakIndexInMountainArray(arr))
}
