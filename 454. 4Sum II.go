package main
//1.将C+D的每一种可能放入查找表
//2.双重循环遍历A B，去C+D的查找表中匹配
func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int)
	for i := 0;i<len(C);i++{
		for j := 0;j<len(D);j++{
			m[C[i]+D[j]]++
		}
	}
	res := 0
	for i := 0;i<len(A);i++{
		for j := 0;j<len(B);j++{
			target := 0-A[i]-B[j]
			if _,ok := m[target];ok{
				res += m[target]
			}
		}
	}
	return res
}