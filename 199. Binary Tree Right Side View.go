package main

import "container/list"

//类似于637号问题
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	l := list.New()
	l.PushBack(root) //首先放入根节点
	tmp := 0
	for l.Len() > 0 {
		size := l.Len() //这个size记录的是当前层的节点个数
		for i := 0; i < size; i++ {
			node := l.Front().Value.(*TreeNode)
			l.Remove(l.Front())
			tmp = node.Val
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
