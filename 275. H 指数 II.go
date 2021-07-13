package main

func hIndex2(citations []int) int {
	l := 0
	n := len(citations)
	r := n
	for l < r {
		mid := l + (r-l)/2
		if citations[mid] >= n-mid {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return n - l
}
