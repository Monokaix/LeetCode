package main

import "container/list"

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	// reverse odd slice
	for index := 0;index<len(res);index++{
		if index % 2 == 1{
			for i,j := 0,len(res[index])-1;i<j;i,j = i+1,j-1{
				res[index][i],res[index][j] = res[index][j],res[index][i]
			}
		}
	}
	return res
}
