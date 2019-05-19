package main

import "fmt"

func countSmaller(nums []int) []int {
	n := len(nums)
	count := make([]int, n)
	arr := make([][]int, 0)
	for i := 0; i < n; i++ {
		tmp := make([]int, 2)
		tmp[0] = nums[i]
		tmp[1] = i
		arr = append(arr, tmp)
	}
	temp := make([][]int, 0)
	for i := 0; i < n; i++ {
		tmp := make([]int, 2)
		temp = append(temp, tmp)
	}
	mergeSort(arr, 0, n-1, temp, &count, nums)
	return count
}

func mergeSort(arr [][]int, left, right int, temp [][]int, count *[]int, nums []int) {
	if left < right {
		mid := (right + left) / 2
		mergeSort(arr, left, mid, temp, count, nums)
		mergeSort(arr, mid+1, right, temp, count, nums)
		//merge2(arr, left, mid, right, temp, count)
		//合并两个有序数组
		i := left
		j := mid + 1
		t := 0 //临时数组索引
		for i <= mid && j <= right {
			if arr[i][0] <= arr[j][0] { //谁小谁就放在temp数组最前面，直到两个数组其中一个遍历完，此时跳出循环
				temp[t][0] = arr[i][0]
				temp[t][1] = arr[i][1]
				(*count)[arr[i][1]] += j - mid - 1
				t++
				i++
			} else {
				temp[t][0] = arr[j][0]
				temp[t][1] = arr[j][1]
				t++
				j++
			}
		}
		//将左边序列元素放入temp
		for i <= mid {
			temp[t][0] = arr[i][0]
			temp[t][1] = arr[i][1]
			(*count)[arr[i][1]] += j - mid - 1
			t++
			i++
		}
		//将右边序列元素放入temp
		for j <= right {
			temp[t][0] = arr[j][0]
			temp[t][1] = arr[j][1]
			t++
			j++
		}
		//temp数组整体拷贝至原数组
		t = 0
		for left <= right {
			arr[left][0] = temp[t][0]
			arr[left][1] = temp[t][1]
			left++
			t++
		}
	}
}

func main() {
	nums := []int{2, 0, 1}
	fmt.Println(countSmaller(nums))
}
