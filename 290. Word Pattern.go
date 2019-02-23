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
	for i := 1; i < len(pattern); i++ {
		if _, ok := m[pattern[i]]; ok {
			if m[pattern[i]] != s[i] {

				return false
			}
		} else {
			m2[s[i-1]] = true
			if _, ok := m2[s[i]]; ok {
				return false
			} else {
				m[pattern[i]] = s[i]
			}

		}

	}
	return true
}

func main() {

}
