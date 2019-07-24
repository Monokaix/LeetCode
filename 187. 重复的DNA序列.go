package main

//因为题目要找的是10个字符,是常数,所以直接遍历一遍取10个做映射进行统计就行
func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	m := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		m[s[i:i+10]]++
	}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
