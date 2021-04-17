package main

import (
	"sort"
)

func isScramble(s1 string, s2 string) bool {
	m := make(map[string]bool)

	var recursive func(t1, t2 []byte, m map[string]bool) bool
	recursive = func(t1, t2 []byte, m map[string]bool) bool {
		n := len(t1)
		if string(t1) == string(t2) {
			return true
		}

		if v, ok := m[string(t1)+"/"+string(t2)]; ok {
			if v {
				return true
			}
			return false
		}

		s1 := make([]byte, n)
		s2 := make([]byte, n)
		copy(s1, t1)
		copy(s2, t2)
		sort.Slice(s1, func(i, j int) bool {
			return s1[i] < s1[j]
		})
		sort.Slice(s2, func(i, j int) bool {
			return s2[i] < s2[j]
		})

		if string(s1) != string(s2) {
			return false
		}

		for i := 1; i < n; i++ {
			if recursive(t1[:i], t2[:i], m) && recursive(t1[i:], t2[i:], m) || recursive(t1[:i], t2[n-i:], m) && recursive(t1[i:], t2[:n-i], m) {
				m[string(t1)+"/"+string(t2)] = true
				return true
			}
		}
		m[string(t1)+"/"+string(t2)] = false
		return false
	}

	return recursive([]byte(s1), []byte(s2), m)
}
