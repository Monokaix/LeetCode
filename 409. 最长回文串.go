package main

func longestPalindrome2(s string) int {
	m := [128]int{}
	flag := 0
	res := 0
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < 128; i++ {
		if m[i]%2 == 0 {
			res += m[i]
		} else {
			res += m[i] - 1
			flag = 1
		}
	}
	return res + flag
}
