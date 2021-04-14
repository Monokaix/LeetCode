package main

import (
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	i := 2
	preStr := "1"
	for i <= n {
		j := 0
		tmp := ""
		preIndex := 0
		for j < len(preStr) {
			if j == len(preStr)-1 || preStr[j] != preStr[j+1] {
				tmp += strconv.Itoa(j-preIndex+1) + string(preStr[j])
				preIndex = j + 1
			}
			j++
		}

		preStr = tmp
		i++
	}

	return preStr
}
