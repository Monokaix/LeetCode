package main

func reversePairs(nums []int) int {
	res := 0
	tmp := make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1, tmp, &res)
	return res
}

func mergeSort(nums []int, left, right int, tmp []int, res *int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(nums, left, mid, tmp, res)
	mergeSort(nums, mid+1, right, tmp, res)
	merge(nums, left, mid, right, tmp, res)
}

func merge(nums []int, left, mid, right int, tmp []int, res *int) {
	i := left
	j := mid + 1
	t := 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[t] = nums[i]
			i++
			*res += j - mid - 1
		} else {
			tmp[t] = nums[j]
			j++
		}
		t++
	}

	for i <= mid {
		*res += j - mid - 1
		tmp[t] = nums[i]
		t++
		i++
	}
	for j <= right {
		tmp[t] = nums[j]
		t++
		j++
	}
	t = 0
	for i := left; i <= right; i++ {
		nums[i] = tmp[t]
		t++
	}
}

//func main() {
//	nums := []int{7, 5, 6, 4}
//	reversePairs(nums)
//	fmt.Println(nums)
//}
