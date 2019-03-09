package main

import (
	"strings"
	"fmt"
)

func simplifyPath(path string) string {
	str := strings.Split(path, "/")
	res := make([]string, 0) // spilt path without /
	for _,v := range str {
		if v != "" {
			res = append(res, v)
		}
	}
	size := len(res)
	stack := make([]string,size)
	top := 0
	for i:=0;i<size;i++{
		v := res[i]
		if v != "." && v != ".."{
			stack[top] = v
			top++
		}
		if v == "."{
			continue
		}
		if v == ".."{
			if top >0{
				top--
			}
		}
	}
	if top <= 0{
		return "/"
	}
	s := ""
	for i:= 0;i<top;i++{ //traverse the stack to return final path
		s += "/"+stack[i]
	}
	return s

}
func main() {
	res := simplifyPath("/a/../../b/../c//.//")
	fmt.Println(res)
}
