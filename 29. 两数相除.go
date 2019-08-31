package main

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	flag := true
	if dividend^divisor < 0 {
		flag = false
	}
	if divisor > 0 {
		divisor = -divisor
	}
	if dividend > 0 {
		dividend = -dividend
	}
	res := 0
	multiple := 1
	tmp_divisor := divisor
	for dividend <= tmp_divisor {
		if dividend <= divisor { //每次除数成倍数递增,减少减法运算次数
			res += multiple
			dividend -= divisor
			divisor += divisor
			multiple += multiple
		} else { //超出时,就是是除数大于被除数重新计算
			divisor = tmp_divisor
			multiple = 1
		}
	}
	if flag {
		return res
	}
	return -res
}
