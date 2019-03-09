package main

import "strconv"

func evalRPN(tokens []string) int {
	size := len(tokens)
	stack := make([]int,size)
	top := 0
	for i := 0;i<size;i++{
		token ,e := strconv.Atoi(tokens[i])
		if e == nil{
			stack[top] = token
			top++
		}else {
			res := 0
			num1:=stack[top-2]
			num2:= stack[top-1]
			switch tokens[i] {
			case "+":
				res = num1+num2
			case "-":
				res = num1-num2
			case "*":
				res = num1*num2
			case "/":
				res = num1/num2
			}
			stack[top-2] = res
			top--
		}
	}
	return stack[0]
}