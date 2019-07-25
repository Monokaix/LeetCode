package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sum := len(nums1) + len(nums2)
	res := 0.0
	if sum%2 == 0 {
		mid1 := findKth(nums1, nums2, sum/2)
		mid2 := findKth(nums1, nums2, sum/2+1)
		res = (mid1 + mid2) / 2
	} else {
		res = findKth(nums1, nums2, sum/2+1)
	}
	return res
}

//思路：使用类二分法,比较A[k/2-1]和B[k/2-1]的大小，相等则返回两者中较大者
//若A[k/2-1]<B[k/2-1],说明A的前k/2-1个元素在前k个最小元素中,因此直接剔除，同理反之则剔除B的
func findKth(A []int, B []int, k int) float64 {
	m := len(A)
	n := len(B)
	//技巧处理,每次都让A最小
	if m > n {
		return findKth(B, A, k)
	}
	if m == 0 { //第k大的数根据B的索引找
		return float64(B[k-1])
	}
	if k == 1 { //k=1时直接比较A和B第一个元素最大值
		return math.Min(float64(A[0]), float64(B[0]))
	}
	//以上是递归终止条件
	a := int(math.Min(float64(k/2), float64(m)))
	b := k - a
	if A[a-1] == B[b-1] { //找到第K大的值,直接返回
		return float64(A[a-1])
	} else if A[a-1] < B[b-1] {
		A = A[a:]
		return findKth(A, B, k-a)
	} else {
		B = B[b:]
		return findKth(A, B, k-b)
	}
}
func main() {
	nums1 := []int{1, 2,3,5,6,89}
	nums2 := []int{0,5,7,8}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
