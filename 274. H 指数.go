package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	res := 0
	n := len(citations)
	for i := 0; i < n; i++ {
		if citations[i] >= n-i {
			if n-i > res {
				res = n - i
			}
		}
	}

	return res
}

func main() {
	nums := []int{3, 0, 6, 1, 5, 7}
	fmt.Println(hIndex(nums))
}
