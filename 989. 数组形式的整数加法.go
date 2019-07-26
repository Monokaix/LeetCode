package main

import "fmt"

func addToArrayForm(A []int, K int) []int {
	m := len(A)
	n := 0
	K2 := K
	for K2 > 0 {
		K2 /= 10
		n++
	}
	length := m
	if m < n {
		length = n
	}
	res := make([]int, length+1)
	carry := 0
	i := m - 1
	for i >= 0 || K > 0 || carry > 0 {
		val1 := 0
		val2 := 0
		if i >= 0 {
			val1 = A[i]
			i--
		}
		if K > 0 {
			fmt.Println(K)
			val2 = K % 10
			K /= 10
		}
		res[length] = (val1 + val2 + carry) % 10
		carry = (val1 + val2 + carry) / 10
		length--
	}
	if res[0] == 0 {
		res = res[1:]
	}
	return res
}
func main() {
	a := []int{0}
	fmt.Println(addToArrayForm(a, 0))
}
