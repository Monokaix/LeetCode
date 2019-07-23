package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		tmp := sortString(strs[i])
		m[tmp] = append(m[tmp], strs[i])
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
func sortString(s string) string {
	res := ""
	tmp := []string{}
	for i := 0; i < len(s); i++ {
		tmp = append(tmp, string(s[i]))
	}
	sort.Strings(tmp)
	for i := 0; i < len(tmp); i++ {
		res += tmp[i]
	}
	return res
}
