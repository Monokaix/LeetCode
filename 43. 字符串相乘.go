package main

import "fmt"

//思路：用字符串模拟整数相乘,两个数的每一位分别相乘,然后存到数组,然后两层循环对保存结果的数组的每一位进行累加
//num1的第i位和num2的第j位相乘得到的是两位数,放在res数组里的i+j和i+j+1的位置上
func multiply(num1 string, num2 string) string {
	m := len(num1)
	n := len(num2)
	if num1 == "0"|| num2 == "0" {
		return "0"
	}
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- { //从后往前计算,因为低位(个位)在后面
		for j := n - 1; j >= 0; j-- {
			tmp := int((num1[i] - '0') * (num2[j] - '0'))
			tmp += res[i+j+1] //累加原来的值
			res[i+j] += tmp / 10
			res[i+j+1] = tmp % 10
		}
	}
	s := ""
	i := 0
	fmt.Println("res", res)
	for i < m+n && res[i] == 0 { //去掉前面的0
		i++
	}
	for i < m+n {
		s += string(res[i] + '0')
		i++
	}
	return s
}

