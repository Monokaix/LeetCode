package main

import "fmt"

func generateParenthesis(n int) []string {
	res := []string{}
	genParenthesis("", n, n, &res)
	return res
}

//left right分别表示还可放置的左右括号的剩余数
func genParenthesis(item string, left int, right int, res *[]string) {
	if left == 0 && right == 0 { //左右括号都放完
		*res = append(*res, item)
		return
	}
	//下面两个是规则，这样保证放入的都是合法序列
	if left > 0 { //先放左括号，因此先判断left是否有剩余
		genParenthesis(item+"(", left-1, right, res)
	}
	//什么时候可以放右括号呢，就是右括号剩余数量比左括号多时
	if left < right {
		genParenthesis(item+")", left, right-1, res)
	}
	//要想满足条件，则递归有限制
	//1、左右括号数量各为n，有一个超过n，则递归终止
	//2、必须先放左括号
	//3、若左括号小于等于右括号数量，则不可放置左括号
}
func main() {
	fmt.Println(generateParenthesis(2))
}
