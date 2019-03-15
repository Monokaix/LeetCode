package main

import (
	"fmt"
)

func partition2(s string) [][]string {
	str := []byte(s)
	res := make([][]string,0)
	backTracking(str,[]string{},&res)
	return res
}

func backTracking(s []byte,temp []string,res *[][]string){
	if len(s) == 0{
		tmp := make([]string,len(temp))
		copy(tmp,temp)

		*res = append(*res,tmp)
		return
	}
	for i:=1;i<=len(s);i++{
		if isPali(string(s[:i])){
			backTracking(s[i:],append(temp, string(s[:i])),res)
		}
	}
}
func isPali(s string) bool{
	i :=0
	j:=len(s)-1
	for i<j{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}
func main(){
	res := make([][]string,0)
	s := []byte("cbbbcc")
	backTracking(s,[]string{},&res)
	fmt.Println(res)
}