package main

func sortArrayByParity(A []int) []int {
	if len(A) < 2 {
		return A
	}
	left := 0
	right := len(A) - 1
	for left < right {
		for left < right && A[left]%2 == 0 { //左边指向的是偶数就++
			left++
		}

		for left < right && A[right]%2 == 1 { //右边指向的是奇数就--
			right--
		}
		if left < right {
			A[left], A[right] = A[right], A[left]
			left++
			right--
		}
	}
	return A
}
