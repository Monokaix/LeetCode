package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	s := make([][]byte, 0)
	for i := 0; i < len(strs); i++ {
		tmp := ([]byte)(strs[i])
		s = append(s, tmp)
	}

	minLen := 200
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	cnt := 0
	for j := 0; j < minLen; j++ {
		pre := s[0][j]
		for i := 0; i < len(strs); i++ {
			if s[i][j] != pre {
				return string(s[0][:cnt])
			}
		}
		cnt++
	}

	return string(s[0][:cnt])
}
