package main

import "container/list"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		size := l.Len()
		eachLevel := make([]int, 0)
		for i := 0; i < size; i++ {
			// 每次取出队首元素
			first := l.Front().Value.(*TreeNode)
			eachLevel = append(eachLevel, first.Val)
			l.Remove(l.Front())
			if first.Left != nil {
				l.PushBack(first.Left)
			}
			if first.Right != nil {
				l.PushBack(first.Right)
			}
		}
		// 一层遍历完成
		res = append(res, eachLevel)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i] = res[len(res)-i], res[i]
	}
	return res
}
