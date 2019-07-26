package main

import "fmt"

//模拟每一位相加存在int数组里,注意进位
func addStrings(num1 string, num2 string) string {
	m := len(num1)
	n := len(num2)
	s := ""
	length := m
	if m < n {
		length = n
	}
	res := make([]int, length+1)
	carry := 0
	i := m - 1
	j := n - 1
	for i >= 0 || j >= 0 || carry > 0 {
		val1 := 0
		val2 := 0
		if i >= 0 {
			val1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			val2 = int(num2[j] - '0')
			j--
		}
		res[length] = (val1 + val2 + carry) % 10
		carry = (val1 + val2 + carry) / 10
		length--
	}
	//fmt.Println(res)
	flag := 0
	if res[0] == 0 {
		flag = 1
	}
	for i := flag; i < len(res); i++ {
		s += string(res[i] + '0')
	}
	return s
}

func main() {
	fmt.Println(addStrings("1", "9"))
}
