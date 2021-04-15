package main

import "fmt"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} // 从大到小排序
	arr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			res += arr[i]
			num -= values[i]
		}
	}

	return res
}

func main() {
	nums := 9876
	fmt.Println(intToRoman(nums))
}
