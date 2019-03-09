package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	type Command struct {
		s    string
		node *TreeNode
	}
	res := make([]int, 0)
	stack := make([]Command, 0)
	top := 0
	stack = append(stack, Command{"go", root})
	top++
	for top > 0 {
		command := stack[top-1] // pop stack every time
		stack = stack[:top-1]
		top--
		if command.s == "print" {
			res = append(res, command.node.Val)
		} else {
			if command.node.Right != nil {
				stack = append(stack, Command{"go", command.node.Right})
				top++
			}
			if command.node.Left != nil {
				stack = append(stack, Command{"go", command.node.Left})
				top++
			}
			stack = append(stack, Command{"print", command.node})
			top++
		}
	}
	return res
}
func main(){

	node1 := &TreeNode{
		2,
		nil,
		nil,
	}
	node2 := &TreeNode{
		3,
		nil,
		nil,
	}
	root := &TreeNode{
		1,
		node1,
		node2,
	}
	fmt.Println(preorderTraversal(root))
}