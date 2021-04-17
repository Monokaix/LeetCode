package main

import (
	"fmt"
	"sort"
)

type Pair4 struct {
	num   int
	index int
}

func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	p := make([]Pair4, n)
	for i := 0; i < n; i++ {
		p[i] = Pair4{nums[i], i}
	}

	sort.Slice(p, func(i, j int) bool {
		if p[i].num == p[j].num {
			return p[i].index < p[j].index
		}

		return p[i].num < p[j].num
	})

	abs := func(a, b int) int {
		if a > b {
			return a - b
		}

		return b - a
	}

	for i := 0; i < n-1; i++ {
		j := i + 1
		for j < n {
			if p[j].num-p[i].num > t {
				break
			}

			if abs(p[j].index, p[i].index) <= k {
				fmt.Println(p[j], p[i])
				return true
			}
			j++
		}
	}

	return false
}

func main() {
	nums := []int{0, 10, 22, 15, 0, 5, 22, 12, 1, 5}
	containsNearbyAlmostDuplicate2(nums, 3, 3)
}
