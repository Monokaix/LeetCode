package main

import "fmt"

// 贪心算法，因为暴力穷举复杂度太高不现实
// 利用一个栈，遍历字符串的每一个值时比较当前元素与栈顶元素大小，
// 若小于栈顶，则栈顶出栈，当前元素进栈，因为这样才能保证去掉的元素最大，从而使得
// 剩下字符组成的数字最小
// 注意问题
// 1、弹出栈顶元素的条件：栈不为空&&k>0&&当前元素小于栈顶元素
// 2、若遍历完字符串，k仍然大于0，则需要从栈顶继续删除元素
// 3、字符串中含有0，要做特殊处理
func removeKdigits(num string, k int) string {
	stack := []byte{}
	top := 0
	for i := 0; i < len(num); i++ {
		for top > 0 && k > 0 && num[i] < stack[top-1] { //满足条件则出栈
			stack = stack[:top-1]
			top--
			k--
		}
		if num[i] != '0' || top > 0 { // 这里栈为空且当前元素为0不入栈，如num=10020，k=1
			stack = append(stack, num[i]) //不满足条件，直接入栈
			top++
		}
	}
	// 如果k仍大于0，从末尾开始删除
	for k > 0 && top > 0 {
		stack = stack[:top-1]
		top--
		k--
	}
	if top ==0{
		return "0"
	}
	return string(stack)
}
func main() {
	s := "1432219"
	fmt.Println(removeKdigits(s, 3))
}
