package main

func postorderTraversal(root *TreeNode) []int {
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
			stack = append(stack, Command{"print", command.node})
			top++
			if command.node.Right != nil {
				stack = append(stack, Command{"go", command.node.Right})
				top++
			}
			if command.node.Left != nil {
				stack = append(stack, Command{"go", command.node.Left})
				top++
			}

		}
	}
	return res
}
