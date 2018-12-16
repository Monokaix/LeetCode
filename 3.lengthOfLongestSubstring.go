package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	freq :=[256]int{}
	l := 0
	r :=-1
	res := 0
	for l < len(s){
		if r+1 < len(s) && freq[s[r+1]] == 0{ //判断下一个字符是否已经存在于当前的滑动窗口中 r+1
			r++
			freq[s[r]]++
		}else{
			freq[s[l]]-- //抛弃左边重复的元素
			l++ //左边指针向前移动
		}
		if res < (r-l+1){
			res = r-l+1
		}
	}
	return res
}

func main() {
	s := "abcbbccc"
	fmt.Println(lengthOfLongestSubstring(s))
}
