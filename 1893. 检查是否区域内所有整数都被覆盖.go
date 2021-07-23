package main

func isCovered(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		cover := false
		for j := range ranges {
			if i >= ranges[j][0] && i <= ranges[j][1] {
				cover = true
				break
			}
		}
		if !cover {
			return false
		}
	}

	return true
}

func isCovered2(ranges [][]int, left, right int) bool {
	diff := [52]int{} // 差分数组
	for _, r := range ranges {
		diff[r[0]]++
		diff[r[1]+1]--
	}
	sum := [52]int{} // 前缀和
	for i := 1; i < 52; i++ {
		sum[i] += sum[i-1] + diff[i]
	}
	for i := left; i <= right; i++ {
		if sum[i] <= 0 {
			return false
		}
	}
	return true
}

func main() {
	ranges := [][]int{{1, 2}, {3, 4}, {5, 6}}
	isCovered2(ranges, 2, 5)
}
