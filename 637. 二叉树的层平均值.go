package main

import "container/list"

//一种方法是每一次都存储下来,层序遍历完计算每层的平均值,但该方法占用空间较大
//因此考虑遍历每层的时候就进行计算,不单独进行存储
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	l := list.New()
	l.PushBack(root) //首先放入根节点
	for l.Len() > 0 {
		size := l.Len() //这个size记录的是当前层的节点个数
		sum := 0.0
		for i := 0; i < size; i++ {
			node := l.Front().Value.(*TreeNode)
			l.Remove(l.Front())
			sum += float64(node.Val)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		res = append(res, sum/float64(size))
	}
	return res
}
