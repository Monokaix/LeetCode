package main

//只需关注偶数位置上放的全是偶数即可
func sortArrayByParityII(A []int) []int {
	j := 1 //i指向偶数位置,j指向奇数位置
	for i := 0; i < len(A); i = i + 2 {
		if A[i]&1 == 1 { //如果偶数位置上的数是奇数,就一直往下找j的位置,直到找到一个偶数
			for A[j]&1 == 1 {
				j += 2
			}
			//找到一个偶数,交换
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
