package main

import (
	"fmt"
	"math"
)

//思路：充分利用满二叉树的性质
//在第level层时,当前节点值为label,则可以找到其对称节点为b=2^level-1+2^(level-1)-label
//2^level-1+2^(level-1)即为该层对称节点总和
//找到其对称节点后,该对称节点的父节点即为当前label的上一层父节点
//然后每次循环直到label为0
func pathInZigZagTree(label int) []int {
	level := int(math.Log2(float64(label)))
	res := make([]int, level+1)
	fmt.Println(level)
	for level > 0 { //循环直到到达根节点
		//存放子节点的label
		res[level] = label
		//值得到对称节点
		label = int(math.Pow(2, float64(level))) + int(math.Pow(2, float64(level+1))) - 1 - label
		//除以2,得到的即为当前之字形树下label的父节点
		label >>= 1
		level--
	}
	return res
}