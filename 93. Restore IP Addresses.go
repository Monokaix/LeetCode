package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	nums := []byte(s)
	dfs(nums, 1, "", &res)
	return res
}

func dfs(nums []byte, num int, temp string, res *[]string) {
	if len(nums) == 0 {
		return
	}
	if num == 4 {
		if isLegal(string(nums)) {
			temp += string(nums)
			*res = append(*res, temp)
			return
		}
		return
	}
	for i := 1; i < 4 && i < len(nums); i++ {
		if isLegal(string(nums[:i])) {
			dfs(nums[i:], num+1, temp+string(nums[:i])+".", res)
		} else {
			return
		}
	}
	return
}

func isLegal(s string) bool {
	num, _ := strconv.Atoi(s)
	if num <= 255 && strconv.Itoa(num) == s {
		return true
	}
	return false
}
func main() {
	res := make([]string, 0)
	nums := []byte("010010")
	//fmt.Println(nums[3:])
	dfs(nums, 1, "", &res)
	fmt.Println(res)
	//fmt.Println(len(res))

}
