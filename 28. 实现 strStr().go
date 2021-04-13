package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	m := len(haystack)
	n := len(needle)
	i := 0
	j := 0

	for i < m {
		for i < m && haystack[i] != needle[j] {
			i++
		}

		// 遍历完第一个字符串没有匹配的，直接返回
		if i >= m {
			return -1
		}

		// index记录上次比较的位置，可能是结果
		index := i
		for i < m && j < n && haystack[i] == needle[j] {
			i++
			j++
		}

		// 说明有不相等的，此时需要i往前移动重新比较
		if j < n {
			j = 0
			i = index + 1
		} else {
			return index
		}
	}

	return -1
}

func main() {
	a := "aaaa"
	b := "aaa"
	fmt.Println(strStr(a, b))
}
