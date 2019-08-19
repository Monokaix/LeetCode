package main

import "container/list"

//func levelOrder(root *TreeNode) [][]int {
//	res := [][]int{}
//	if root == nil {
//		return res
//	}
//	l := list.New()
//	type info struct {
//		node  *TreeNode
//		level int
//	}
//	l.PushBack(info{root, 0})
//	for l.Len() > 0 {
//		node := l.Front().Value.(info).node
//		level := l.Front().Value.(info).level
//		l.Remove(l.Front())
//		// new level
//		if level >= len(res) {
//			res = append(res, []int{})
//		}
//		res[level] = append(res[level], node.Val)
//		// left and right node to queue
//		if node.Left != nil {
//			l.PushBack(info{node.Left, level + 1})
//		}
//		if node.Right != nil {
//			l.PushBack(info{node.Right, level + 1})
//		}
//	}
//	return res
//}
//和637号问题类似
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	l := list.New()
	l.PushBack(root) //首先放入根节点
	for l.Len() > 0 {
		size := l.Len() //这个size记录的是当前层的节点个数
		each := make([]int, 0)
		for i := 0; i < size; i++ {
			node := l.Front().Value.(*TreeNode)
			l.Remove(l.Front())
			each = append(each, node.Val)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		//一层遍历完成
		res = append(res, each)
	}
	return res
}
