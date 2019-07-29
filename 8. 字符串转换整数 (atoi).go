package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	flag := 1 //代表正数
	j := 0
	for j < len(str) {
		if str[j] == ' ' {
			j++ //是空格就跳过
		} else {
			break
		}
	}
	if str[j:] == "" { //去掉开头空格后字符串为空
		return 0
	}
	if str[j] == '-' {
		flag = -1 //负数
		j++
	} else if str[j] == '+' {
		j++
	}
	num := 0
	for i := j; i < len(str); i++ {

		if str[i] >= '0' && str[i] <= '9' {
			cur := int(str[i]-'0') * flag
			if num > math.MaxInt32/10 || num == math.MaxInt32/10 && cur > 7 { //正数溢出
				return math.MaxInt32
			}
			if num < math.MinInt32/10 || num == math.MinInt32/10 && cur < -8 { //负数溢出
				return math.MinInt32
			}
			num = num*10 + cur
		} else {
			if j == i { //第一个有效字符不合法
				return 0
			}
			break
		}
	}
	return num
}
func main() {
	str := "2147483646"
	fmt.Println(myAtoi(str))
}
