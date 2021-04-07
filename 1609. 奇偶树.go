package main

import (
	"container/list"
	"math"
)

type treePair struct {
	node  *TreeNode
	level int
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := list.New()
	l.PushBack(&treePair{
		node:  root,
		level: 0,
	})

	for l.Len() > 0 {
		size := l.Len()
		preMin := math.MinInt32
		preMax := math.MaxInt32

		for i := 0; i < size; i++ {
			p := l.Front().Value.(*treePair)
			l.Remove(l.Front())
			if p.level%2 == 0 {
				if p.node.Val%2 == 0 {
					return false
				}

				if p.node.Val <= preMin {
					return false
				}
				preMin = p.node.Val
			} else {
				if p.node.Val%2 == 1 {
					return false
				}

				if p.node.Val >= preMax {
					return false
				}
				preMax = p.node.Val
			}

			if p.node.Left != nil {
				l.PushBack(&treePair{
					node:  p.node.Left,
					level: p.level + 1,
				})
			}
			if p.node.Right != nil {
				l.PushBack(&treePair{
					node:  p.node.Right,
					level: p.level + 1,
				})
			}
		}
	}

	return true
}
