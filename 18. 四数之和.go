package main

import (
	"sort"
)

//func fourSum(nums []int, target int) [][]int {
//	res := [][]int{}
//	cur := []int{}
//	m := make(map[string]bool)
//	dfs6(nums, 0, target, cur, &res, &m)
//	return res
//}
//
//func dfs6(nums []int, begin, target int, cur []int, res *[][]int, m *map[string]bool) {
//	if len(cur) > 4 {
//		return
//	}
//	if len(cur) == 4 {
//		tmp := 0
//		for i := 0; i < 4; i++ {
//			tmp += cur[i]
//		}
//
//		if tmp == target {
//			ans := make([]int, len(cur))
//			copy(ans, cur)
//			sort.Ints(ans)
//			key := ""
//			for i := 0; i < 4; i++ {
//				key += strconv.Itoa(ans[i])
//			}
//			if _, ok := (*m)[key]; !ok {
//				*res = append(*res, ans)
//				(*m)[key] = true
//			}
//		}
//
//		return
//	}
//
//	for i := begin; i < len(nums); i++ {
//		cur = append(cur, nums[i])
//		dfs6(nums, i+1, target, cur, res, m)
//		cur = cur[:len(cur)-1]
//	}
//}

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	res := [][]int{}
	if n < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		// 为了去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 剪枝
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			return res
		}
		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// 去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 剪枝
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}

			left := j + 1
			right := n - 1

			for left < right {
				cur := nums[i] + nums[j] + nums[left] + nums[right]
				if cur == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < n && nums[left] == nums[left-1] {
						left++
					}
					right--
					for right > -1 && nums[right] == nums[right+1] {
						right--
					}
				} else if cur < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}
