package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	str := strings.Split(s, " ")
	res := ""
	for i := 0; i < len(str); i++ {
		if i != len(str)-1 {
			res += reverseStr(str[i]) + " "
		} else {
			res += reverseStr(str[i])
		}
	}
	return res
}

func reverseStr(str string) string {
	tmp := []rune(str)
	for i := 0; i < len(tmp)/2; i++ {
		tmp[i], tmp[len(str)-i-1] = tmp[len(str)-i-1], tmp[i]
	}
	return string(tmp)
}

func main() {
	str := "abc"
	fmt.Println(reverseStr(str))
}
