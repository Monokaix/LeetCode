package main

import "strings"

func wordPattern(pattern string, str string) bool {
	m := make(map[byte]string)
	m2 := make(map[string]bool)
	s := strings.Split(str, " ")
	if len(pattern) != len(s) {
		return false
	}
	m[pattern[0]] = s[0]
	m2[s[0]] = true
	for i := 1; i < len(pattern); i++ {
		if _, ok := m[pattern[i]]; ok { //出现重复pattern字符时,看其映射是否正确
			if m[pattern[i]] != s[i] {
				return false
			}
		} else { //当patterns中出现一个新字符时,要看其对应位置的s串是否也出现过

			if _, ok := m2[s[i]]; ok { //若已经出现,直接返回false
				return false
			} else { //若未出现,建立pattern到s的映射
				m[pattern[i]] = s[i]
				m2[s[i]] = true
			}
		}
	}
	return true
}
