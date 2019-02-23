package main

//1.i是一个枢纽，对于每一个i遍历他到其他点的距离
//2.用一张查找表（map）记录i到其他点的距离，map key:距离，value:距离相等的点的个数
//3.根据map中的value计算总个数，value*(value-1)

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0;i<len(points);i++{
		m := make(map[int]int)
		for j := 0;j<len(points);j++{
			if i != j{
				m[dist(points[i],points[j])]++
			}
		}
		for k := range m{
			res += m[k]*(m[k]-1)
		}
	}
	return res

}

func dist(num1 []int,num2 []int)int{
	return (num1[0]-num2[0]) * (num1[0]-num2[0])+
		(num1[1]-num2[1])*(num1[1]-num2[1])
}