package main

var letterMap []string

func letterCombinations(digits string) []string {
	res := make([]string,0)
	if digits == ""{
		return res
	}
	letterMap = []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	findConbination(digits,0,"",&res)
	return res
}

func findConbination(digits string, index int, s string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, s)
		return
	}
	c := digits[index]
	letters := letterMap[c-'0']
	for i := 0; i < len(letters); i++ {
		findConbination(digits, index+1, s+string(letters[i]),res)
	}
}
