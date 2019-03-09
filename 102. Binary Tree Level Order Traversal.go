package main

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	l := list.New()
	type info struct {
		node  *TreeNode
		level int
	}
	l.PushBack(info{root, 0})
	for l.Len() > 0 {
		node := l.Front().Value.(info).node
		level := l.Front().Value.(info).level
		l.Remove(l.Front())
		// new level
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)
		// left and right node to queue
		if node.Left != nil {
			l.PushBack(info{node.Left, level + 1})
		}
		if node.Right != nil {
			l.PushBack(info{node.Right, level + 1})
		}
	}
	return res
}
